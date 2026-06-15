// Go Interview Guide - Shared Scripts

const SECTION_TOTALS = {
  fundamentals: 8,
  concurrency: 8,
  'data-structures': 8,
  algorithms: 12,
  'systems-design': 8,
  projects: 8
};

let workspaceKeys = new Set();
let collapsedKeys = new Set();

document.addEventListener('DOMContentLoaded', function() {
  initCopyButtons();
  initToggleButtons();
  initActiveNav();
  initTemplateLoading();
  initGenerateButtons();
  initWorkspaceState();
});

// Copy buttons
function initCopyButtons() {
  document.querySelectorAll('.copy-btn').forEach(btn => {
    btn.addEventListener('click', async function() {
      const wrapper = this.closest('.code-block-wrapper');
      const activePane = wrapper.querySelector('.tab-pane.active');
      const text = activePane ? activePane.textContent : '';
      try {
        await navigator.clipboard.writeText(text);
        const originalText = this.innerHTML;
        this.innerHTML = '<svg width="12" height="12" viewBox="0 0 16 16" fill="currentColor"><path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.75.75 0 0 1 1.06-1.06L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0z"/></svg> Copied!';
        this.classList.add('copied');
        setTimeout(() => {
          this.innerHTML = originalText;
          this.classList.remove('copied');
        }, 2000);
      } catch (err) {
        // Fallback
        const textarea = document.createElement('textarea');
        textarea.value = text;
        textarea.style.position = 'fixed';
        textarea.style.opacity = '0';
        document.body.appendChild(textarea);
        textarea.select();
        document.execCommand('copy');
        document.body.removeChild(textarea);
        const originalText = this.innerHTML;
        this.innerHTML = '<svg width="12" height="12" viewBox="0 0 16 16" fill="currentColor"><path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.75.75 0 0 1 1.06-1.06L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0z"/></svg> Copied!';
        this.classList.add('copied');
        setTimeout(() => {
          this.innerHTML = originalText;
          this.classList.remove('copied');
        }, 2000);
      }
    });
  });
}

// Toggle buttons (hints, follow-ups, context)
function initToggleButtons() {
  document.querySelectorAll('.toggle-btn').forEach(btn => {
    btn.addEventListener('click', function() {
      const targetId = this.getAttribute('data-target');
      const target = document.getElementById(targetId);
      if (!target) return;
      const isVisible = target.classList.contains('visible');
      // Close all siblings in the same card
      const card = this.closest('.question-card');
      if (card) {
        card.querySelectorAll('.toggle-content').forEach(el => el.classList.remove('visible'));
        card.querySelectorAll('.toggle-btn').forEach(el => el.classList.remove('active'));
      }
      if (!isVisible) {
        target.classList.add('visible');
        this.classList.add('active');
      }
    });
  });
}

// Active nav link
function initActiveNav() {
  const path = window.location.pathname;
  const filename = path.split('/').pop() || 'index.html';
  document.querySelectorAll('.nav-links a').forEach(link => {
    if (link.getAttribute('href') === filename) {
      link.classList.add('active');
    }
  });
}

// Template loading: fetch all template files and render them as tabs
function initTemplateLoading() {
  document.querySelectorAll('.code-block-wrapper[data-template]').forEach(wrapper => {
    const templateDir = wrapper.getAttribute('data-template');
    const contentEl = wrapper.querySelector('.template-content');
    if (!contentEl || !templateDir) return;

    fetch('/api/template-dir?dir=' + encodeURIComponent(templateDir))
      .then(r => {
        if (!r.ok) throw new Error('Failed to load template directory');
        return r.json();
      })
      .then(data => {
        renderTemplateTabs(wrapper, data.files || []);
      })
      .catch(err => {
        contentEl.textContent = '// Error loading template: ' + err.message;
      });
  });
}

function renderTemplateTabs(wrapper, files) {
  const contentEl = wrapper.querySelector('.template-content');
  if (!contentEl) return;

  if (files.length === 0) {
    contentEl.innerHTML = '<div class="no-template-files">No code files available for this question.</div>';
    return;
  }

  const tabBar = document.createElement('div');
  tabBar.className = 'tab-bar';

  const panesContainer = document.createElement('div');
  panesContainer.className = 'tab-panes';

  files.forEach((file, index) => {
    const isActive = index === 0;

    const tab = document.createElement('button');
    tab.className = 'tab-btn' + (isActive ? ' active' : '');
    tab.textContent = file.name;
    tab.setAttribute('data-filename', file.name);
    tabBar.appendChild(tab);

    const pane = document.createElement('pre');
    pane.className = 'tab-pane' + (isActive ? ' active' : '');
    pane.setAttribute('data-filename', file.name);
    const code = document.createElement('code');
    code.textContent = file.content;
    pane.appendChild(code);
    panesContainer.appendChild(pane);

    tab.addEventListener('click', () => {
      wrapper.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
      wrapper.querySelectorAll('.tab-pane').forEach(p => p.classList.remove('active'));
      tab.classList.add('active');
      pane.classList.add('active');
    });
  });

  contentEl.innerHTML = '';
  contentEl.appendChild(tabBar);
  contentEl.appendChild(panesContainer);
}

// Generate workspace buttons
function initGenerateButtons() {
  document.querySelectorAll('.generate-btn').forEach(btn => {
    btn.addEventListener('click', function() {
      const template = this.getAttribute('data-template');
      const section = this.getAttribute('data-section');
      const problem = this.getAttribute('data-problem');
      const statusEl = document.getElementById('ws-' + problem);

      generateWorkspace(template, section, problem, false, statusEl);
    });
  });
}

async function initWorkspaceState() {
  try {
    const res = await fetch('/api/workspaces');
    if (!res.ok) throw new Error('Failed to load workspaces');
    const keys = await res.json();
    workspaceKeys = new Set(keys);
  } catch (err) {
    workspaceKeys = new Set();
  }

  renderWorkspaceBadges();
  initCollapseButton();
  initHeaderToggle();
  initIndexProgress();
}

function renderWorkspaceBadges() {
  document.querySelectorAll('.question-card').forEach(card => {
    const btn = card.querySelector('.generate-btn');
    if (!btn) return;

    const section = btn.getAttribute('data-section');
    const problem = btn.getAttribute('data-problem');
    const key = section + '-' + problem;
    card.setAttribute('data-workspace-key', key);

    if (workspaceKeys.has(key)) {
      markWorkspaceExists(card);
    }
  });
}

function markWorkspaceExists(card) {
  card.classList.add('has-workspace');
  const header = card.querySelector('.question-header');
  if (header && !header.querySelector('.workspace-indicator')) {
    const indicator = document.createElement('span');
    indicator.className = 'workspace-indicator';
    indicator.textContent = 'Workspace';
    header.appendChild(indicator);
  }
}

function initHeaderToggle() {
  document.addEventListener('click', function(e) {
    const header = e.target.closest('.question-card.has-workspace .question-header');
    if (!header) return;
    const card = header.closest('.question-card');
    if (!card) return;

    if (card.classList.contains('collapsed')) {
      expandCard(card);
    } else {
      collapseCard(card);
    }
  });
}

function collapseCard(card) {
  const key = card.getAttribute('data-workspace-key');
  if (key) collapsedKeys.add(key);
  card.classList.add('collapsed');
  updateCollapseButtonText();
}

function initCollapseButton() {
  const sectionHeader = document.querySelector('.section-header');
  if (!sectionHeader || sectionHeader.querySelector('.collapse-btn')) return;

  const btn = document.createElement('button');
  btn.className = 'collapse-btn';
  btn.textContent = 'Collapse all with workspaces';
  btn.addEventListener('click', toggleCollapseAll);
  sectionHeader.appendChild(btn);
}

function toggleCollapseAll() {
  const btn = document.querySelector('.section-header .collapse-btn');
  if (!btn) return;

  if (collapsedKeys.size === 0) {
    document.querySelectorAll('.question-card.has-workspace').forEach(card => {
      const key = card.getAttribute('data-workspace-key');
      if (!key) return;
      collapsedKeys.add(key);
      card.classList.add('collapsed');
    });
    btn.textContent = 'Expand all';
  } else {
    collapsedKeys.clear();
    document.querySelectorAll('.question-card.collapsed').forEach(card => {
      card.classList.remove('collapsed');
    });
    btn.textContent = 'Collapse all with workspaces';
  }
}

function expandCard(card) {
  const key = card.getAttribute('data-workspace-key');
  if (key) collapsedKeys.delete(key);
  card.classList.remove('collapsed');
  updateCollapseButtonText();
}

function updateCollapseButtonText() {
  const btn = document.querySelector('.section-header .collapse-btn');
  if (!btn) return;
  btn.textContent = collapsedKeys.size > 0 ? 'Expand all' : 'Collapse all with workspaces';
}

function initIndexProgress() {
  const cards = document.querySelectorAll('.card-grid .card');
  if (cards.length === 0) return;

  const counts = {};
  workspaceKeys.forEach(key => {
    const lastHyphen = key.lastIndexOf('-');
    const section = lastHyphen > 0 ? key.slice(0, lastHyphen) : key;
    counts[section] = (counts[section] || 0) + 1;
  });

  cards.forEach(card => {
    const section = card.getAttribute('data-section');
    if (!section) return;
    const total = SECTION_TOTALS[section] || 0;
    const started = counts[section] || 0;
    const meta = card.querySelector('.card-meta');
    if (meta) {
      meta.textContent = meta.textContent.replace(/· (Senior & Staff|All Levels|Junior to Staff)/, `· ${started}/${total} started · $1`);
    }
  });
}

function generateWorkspace(template, section, problem, overwrite, statusEl) {
  if (statusEl) {
    statusEl.textContent = 'Generating...';
    statusEl.className = 'workspace-status pending';
  }

  fetch('/api/generate', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ template, section, problem, overwrite })
  })
    .then(r => r.json())
    .then(data => {
      const key = section + '-' + problem;
      const card = statusEl ? statusEl.closest('.question-card') : null;

      if (data.exists && !data.created) {
        // Ask for overwrite confirmation
        showOverwriteModal(template, section, problem, statusEl);
        if (statusEl) {
          statusEl.textContent = 'Workspace already exists at ' + data.path;
          statusEl.className = 'workspace-status exists';
        }
      } else if (data.created) {
        if (statusEl) {
          statusEl.innerHTML = 'Workspace created at <code>' + data.path + '</code>';
          statusEl.className = 'workspace-status created';
        }
      }

      if (card) {
        workspaceKeys.add(key);
        card.setAttribute('data-workspace-key', key);
        markWorkspaceExists(card);
      }
    })
    .catch(err => {
      if (statusEl) {
        statusEl.textContent = 'Error: ' + err.message;
        statusEl.className = 'workspace-status error';
      }
    });
}

function showOverwriteModal(template, section, problem, statusEl) {
  // Remove existing modal if any
  const existing = document.getElementById('overwrite-modal');
  if (existing) existing.remove();

  const modal = document.createElement('div');
  modal.id = 'overwrite-modal';
  modal.className = 'modal-overlay';
  modal.innerHTML = `
    <div class="modal-content">
      <h3>Workspace Already Exists</h3>
      <p>A workspace for <strong>${section}-${problem}</strong> already exists. Overwrite it?</p>
      <div class="modal-actions">
        <button class="modal-btn modal-btn-secondary" id="modal-cancel">Cancel</button>
        <button class="modal-btn modal-btn-primary" id="modal-overwrite">Overwrite</button>
      </div>
    </div>
  `;
  document.body.appendChild(modal);

  document.getElementById('modal-cancel').addEventListener('click', () => {
    modal.remove();
  });

  document.getElementById('modal-overwrite').addEventListener('click', () => {
    modal.remove();
    generateWorkspace(template, section, problem, true, statusEl);
  });

  modal.addEventListener('click', (e) => {
    if (e.target === modal) modal.remove();
  });
}
