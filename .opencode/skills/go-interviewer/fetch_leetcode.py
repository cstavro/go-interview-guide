#!/usr/bin/env python3
"""Fetch a LeetCode problem and generate a local Go scaffold.

Usage:
    python3 fetch_leetcode.py <leetcode-url-or-slug> <output-dir>
"""

import hashlib
import json
import os
import re
import subprocess
import sys
import urllib.request

GRAPHQL_URL = "https://leetcode.com/graphql"


def extract_slug(url):
    m = re.search(r"/problems/([^/]+)", url)
    if m:
        return m.group(1)
    if re.match(r"^[a-z0-9\-]+$", url):
        return url
    raise ValueError(f"Could not extract slug from: {url}")


def fetch_problem(slug):
    query = {
        "query": """
        query questionData($titleSlug: String!) {
                question(titleSlug: $titleSlug) {
                questionId
                questionFrontendId
                title
                titleSlug
                content
                difficulty
                exampleTestcases
                codeSnippets {
                    langSlug
                    code
                }
            }
        }
        """,
        "variables": {"titleSlug": slug},
        "operationName": "questionData",
    }
    data = json.dumps(query).encode("utf-8")
    req = urllib.request.Request(
        GRAPHQL_URL,
        data=data,
        headers={
            "Content-Type": "application/json",
            "User-Agent": "Mozilla/5.0 (compatible; leetcode-fetcher)",
        },
    )
    with urllib.request.urlopen(req, timeout=30) as resp:
        return json.loads(resp.read().decode("utf-8"))


def get_go_snippet(code_snippets):
    for snippet in code_snippets:
        if snippet.get("langSlug") in ("golang", "go"):
            return snippet.get("code", "")
    return ""


def safe_basename(url):
    parsed = urllib.request.urlparse(url)
    basename = os.path.basename(parsed.path)
    if not basename:
        basename = "image"
    h = hashlib.md5(url.encode()).hexdigest()[:8]
    name, ext = os.path.splitext(basename)
    if not ext:
        ext = ".png"
    # Sanitize
    name = re.sub(r"[^a-zA-Z0-9._-]", "_", name)
    return f"{name}_{h}{ext}"


def download_image(url, assets_dir):
    if url.startswith("//"):
        url = "https:" + url
    elif url.startswith("/"):
        url = "https://leetcode.com" + url
    local_name = safe_basename(url)
    local_path = os.path.join(assets_dir, local_name)
    req = urllib.request.Request(
        url, headers={"User-Agent": "Mozilla/5.0 (compatible; leetcode-fetcher)"}
    )
    try:
        with urllib.request.urlopen(req, timeout=30) as resp:
            with open(local_path, "wb") as f:
                f.write(resp.read())
        return "./assets/" + local_name
    except Exception as e:
        print(f"Warning: failed to download {url}: {e}", file=sys.stderr)
        return url


def rewrite_images(content, assets_dir):
    img_tag_pattern = re.compile(r"<img[^>]*>", re.IGNORECASE)
    src_pattern = re.compile(r'src=(["\'])(.*?)\1', re.IGNORECASE)

    def replace_tag(tag):
        def replace_src(m):
            src = m.group(2)
            if src.startswith("data:"):
                return m.group(0)
            new_src = download_image(src, assets_dir)
            return f'src="{new_src}"'

        return src_pattern.sub(replace_src, tag)

    return img_tag_pattern.sub(lambda m: replace_tag(m.group(0)), content)


def get_go_version():
    try:
        out = subprocess.check_output(["go", "version"], text=True)
        m = re.search(r"go(\d+\.\d+)", out)
        if m:
            return m.group(1)
    except Exception:
        pass
    return "1.22"


def extract_func_name(go_code):
    m = re.search(r"func\s+(\w+)\s*\(", go_code)
    return m.group(1) if m else "Solution"


def uncomment_boilerplate_types(code):
    """LeetCode Go boilerplate sometimes places type definitions inside
    /** ... */ comment blocks (e.g. ListNode, TreeNode). This function
    extracts the type definition so the file compiles locally."""
    pattern = re.compile(r"/\*\*(.*?)\*/", re.DOTALL)

    def repl(m):
        block = m.group(1).strip("\n")
        if "type " in block and "struct" in block:
            lines = block.split("\n")
            type_lines = []
            in_type = False
            for line in lines:
                cleaned = re.sub(r"^\s*\*\s?", "", line)
                if cleaned.strip().startswith("type "):
                    in_type = True
                if in_type:
                    type_lines.append(cleaned)
                if in_type and cleaned.strip() == "}":
                    in_type = False
            if type_lines:
                return "\n".join(type_lines)
        return m.group(0)

    return pattern.sub(repl, code)


def add_stub_returns(code):
    """LeetCode Go boilerplate leaves function bodies empty, which does not
    compile locally. This function inserts a minimal zero-value return so the
    file compiles while still matching LeetCode's signature."""
    pattern = re.compile(
        r"(func\s+\w+\s*\([^)]*\)(?:\s*([\[\]\*]*\w+|\([^)]*\)))?\s*\{)(\s*)\}",
        re.DOTALL,
    )

    def repl(m):
        ret_type = m.group(2)
        body_ws = m.group(3)
        if not ret_type:
            return m.group(1) + body_ws + '\t// TODO: implement\n}'
        ret = ret_type.strip()
        if ret in (
            "int", "int8", "int16", "int32", "int64",
            "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
            "float32", "float64", "complex64", "complex128", "byte", "rune",
        ):
            stub = "\treturn 0\n}"
        elif ret == "bool":
            stub = "\treturn false\n}"
        elif ret == "string":
            stub = '\treturn ""\n}'
        else:
            stub = "\treturn nil\n}"
        return m.group(1) + body_ws + stub

    return pattern.sub(repl, code)


def main():
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} <leetcode-url-or-slug> [output-dir]", file=sys.stderr)
        sys.exit(1)

    url = sys.argv[1]
    slug = extract_slug(url)

    print(f"Fetching problem: {slug} ...")
    resp = fetch_problem(slug)
    if "errors" in resp:
        print(f"GraphQL errors: {resp['errors']}", file=sys.stderr)
        sys.exit(1)

    q = resp["data"]["question"]
    problem_id = q.get("questionFrontendId") or q["questionId"]
    title = q["title"]
    content = q["content"]
    difficulty = q.get("difficulty", "Unknown")
    go_code = get_go_snippet(q.get("codeSnippets", []))

    if len(sys.argv) >= 3:
        out_dir = sys.argv[2]
    else:
        out_dir = os.path.join("workspaces", f"leetcode-{problem_id}-{slug}")

    print(f"Problem number: {problem_id}")

    os.makedirs(out_dir, exist_ok=True)
    assets_dir = os.path.join(out_dir, "assets")
    os.makedirs(assets_dir, exist_ok=True)

    content = rewrite_images(content, assets_dir)

    # README.md
    readme_path = os.path.join(out_dir, "README.md")
    with open(readme_path, "w", encoding="utf-8") as f:
        f.write(f"# {title}\n\n")
        f.write(f"`{difficulty}`\n\n")
        f.write("## Description\n\n")
        f.write(content)
        f.write("\n")

    # main.go
    main_go_path = os.path.join(out_dir, "main.go")
    if go_code:
        go_code = uncomment_boilerplate_types(go_code)
        go_code = add_stub_returns(go_code)
        stripped = go_code.strip()
        if not stripped.startswith("package"):
            go_code = "package main\n\n" + go_code
    else:
        go_code = "package main\n\n// TODO: implement solution\n"
    with open(main_go_path, "w", encoding="utf-8") as f:
        f.write(go_code)
        if not go_code.strip().endswith("}"):
            f.write("\n")

    # go.mod
    go_mod_path = os.path.join(out_dir, "go.mod")
    with open(go_mod_path, "w", encoding="utf-8") as f:
        f.write(f"module {slug}\n\n")
        f.write(f"go {get_go_version()}\n")

    # main_test.go (placeholder)
    main_test_path = os.path.join(out_dir, "main_test.go")
    with open(main_test_path, "w", encoding="utf-8") as f:
        f.write("package main\n\n")
        f.write('import "testing"\n\n')
        f.write("func TestSolution(t *testing.T) {\n")
        f.write('\t// TODO: Add tests based on the problem examples and edge cases.\n')
        f.write('\tt.Fatal("TODO: implement tests")\n')
        f.write("}\n")

    # AGENTS.md
    agents_path = os.path.join(out_dir, "AGENTS.md")
    with open(agents_path, "w", encoding="utf-8") as f:
        f.write(f"# Interviewer Agent: {title}\n\n")
        f.write("## Role\n\n")
        f.write(
            "You are a technical interviewer assessing a candidate's solution to this problem. "
            "Your goal is to evaluate their understanding, guide them without giving away the answer, "
            "and provide an honest, constructive review.\n\n"
        )
        f.write("## Problem Overview\n\n")
        f.write(f"Solve the LeetCode problem: **{title}**.\n\n")
        f.write("## Hints (Progressive Disclosure)\n\n")
        f.write(
            "**Hint 1:** Restate the problem in your own words. "
            "What are the inputs, outputs, and constraints?\n\n"
        )
        f.write(
            "**Hint 2:** Consider the time and space complexity requirements. "
            "Are there data structures or algorithms that naturally fit?\n\n"
        )
        f.write(
            "**Hint 3:** Think about edge cases. "
            "What happens with minimal or maximal inputs?\n\n"
        )
        f.write("## Follow-up Questions & Extensions\n\n")
        f.write(
            "**Extension 1:** Can you optimize your solution further? What is the theoretical lower bound?\n\n"
        )
        f.write(
            "**Extension 2:** How would you adapt your solution if the input constraints changed significantly?\n\n"
        )
        f.write("## Solution Review\n\n")
        f.write(
            "Evaluate for correctness, efficiency (time/space complexity), clarity, and testing coverage.\n"
        )
        f.write(
            "Praise specific good decisions and criticize constructively.\n"
        )

    print(f"Generated scaffold in {out_dir}")


if __name__ == "__main__":
    main()
