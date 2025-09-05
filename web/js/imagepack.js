// imagepack.js - client-side logic for the Imagepack editor

document.addEventListener('DOMContentLoaded', () => {
    // Load example pack embedded in the page if present
    let packJsonEl = document.getElementById('imagepack-json');
    let pack = null;
    if (packJsonEl) {
        try {
            pack = JSON.parse(packJsonEl.textContent);
        } catch (e) {
            console.error('Failed to parse imagepack JSON', e);
        }
    }
    if (!pack) return;

    // build preview list
    const previewList = document.querySelector('.imagepack-preview-list');
    const propertiesPanel = document.querySelector('.imagepack-properties-list');
    const statsArea = document.querySelector('.imagepack-bottom');

    function renderPreview() {
        previewList.innerHTML = '';
        pack.content.forEach((item, idx) => {
            const el = document.createElement('div');
            el.className = 'imagepack-preview-image';
            el.dataset.index = idx;

            // create thumbnail (use actual filename when available, otherwise fallback)
            const thumb = document.createElement('img');
            // If you have real thumbnails served, replace this path construction with proper URL
            // For now use favicon as placeholder demo image
            thumb.src = item.thumbnail || '/static/images/favicon.svg';
            thumb.alt = item.filename || item.id || `image-${idx}`;
            thumb.loading = 'lazy';
            thumb.className = 'imagepack-thumb';
            el.appendChild(thumb);

            const label = document.createElement('div');
            label.className = 'imagepack-preview-label';
            label.textContent = item.filename || item.id || `image-${idx}`;
            el.appendChild(label);

            el.addEventListener('click', () => selectImage(idx));
            previewList.appendChild(el);
        });
        updateStats();
    }

    function updateStats() {
        const folders = new Set();
        pack.content.forEach(i => {
            if (i.filename) {
                const parts = i.filename.split('/');
                if (parts.length > 1) {
                    folders.add(parts.slice(0, -1).join('/'));
                }
            }
        });
        const foldersCount = folders.size || 1;
        const imagesCount = pack.content.length;
        const selected = document.querySelector('.imagepack-preview-image.selected');
        const selectedText = selected ? selected.querySelector('.imagepack-preview-label').textContent : 'None';
        statsArea.querySelector('.folders-count').textContent = `Folders: ${foldersCount}`;
        statsArea.querySelector('.images-count').textContent = `Images: ${imagesCount}`;
        statsArea.querySelector('.selected-count').textContent = `Selected: ${selectedText}`;
    }

    function selectImage(idx) {
        document.querySelectorAll('.imagepack-preview-image').forEach(el => el.classList.remove('selected'));
        const el = document.querySelector(`.imagepack-preview-image[data-index='${idx}']`);
        if (el) el.classList.add('selected');
        renderProperties(idx);
        updateStats();
    }

    function renderProperties(idx) {
        const item = pack.content[idx];
        propertiesPanel.innerHTML = '';
        // NSFW toggle
        const nsfwRow = createRow('NSFW', 'nsfw');
        const nsfwToggle = document.createElement('input');
        nsfwToggle.type = 'checkbox';
        nsfwToggle.checked = !!item.nsfw;
        nsfwToggle.addEventListener('change', () => { item.nsfw = nsfwToggle.checked; });
        nsfwRow.appendChild(nsfwToggle);
        propertiesPanel.appendChild(nsfwRow);

        // Title
        const titleRow = createRow('Title', 'title');
        const titleInput = document.createElement('input');
        titleInput.type = 'text';
        titleInput.value = item.title || '';
        titleInput.addEventListener('input', () => { item.title = titleInput.value; });
        titleRow.appendChild(titleInput);
        propertiesPanel.appendChild(titleRow);

        // Helper to create tag array editors
        function createArrayEditor(label, key) {
            const row = createRow(label, key);
            const container = document.createElement('div');
            container.className = 'array-editor';
            const input = document.createElement('input');
            input.type = 'text';
            input.placeholder = 'comma,separated,values';
            input.value = (item[key] || []).join(', ');
            input.addEventListener('change', () => {
                item[key] = input.value.split(',').map(s => s.trim()).filter(Boolean);
            });
            container.appendChild(input);
            row.appendChild(container);
            propertiesPanel.appendChild(row);
        }

        createArrayEditor('Genders', 'genders');
        createArrayEditor('Subraces', 'subraces');
        createArrayEditor('Tags', 'tags');
        createArrayEditor('Backgrounds', 'backgrounds');
        createArrayEditor('Traits', 'traits');

        // Metadata fields
        const meta = item.metadata || {};
        const metaArtist = createRow('Artist', 'meta-artist');
        const artistInput = document.createElement('input');
        artistInput.type = 'text';
        artistInput.value = meta.artist || '';
        artistInput.addEventListener('input', () => {
            item.metadata = item.metadata || {};
            item.metadata.artist = artistInput.value;
        });
        metaArtist.appendChild(artistInput);
        propertiesPanel.appendChild(metaArtist);

        const metaUrl = createRow('URL', 'meta-url');
        const urlInput = document.createElement('input');
        urlInput.type = 'text';
        urlInput.value = meta.url || '';
        urlInput.addEventListener('input', () => {
            item.metadata = item.metadata || {};
            item.metadata.url = urlInput.value;
        });
        metaUrl.appendChild(urlInput);
        propertiesPanel.appendChild(metaUrl);

        const metaLicense = createRow('License', 'meta-license');
        const licenseInput = document.createElement('input');
        licenseInput.type = 'text';
        licenseInput.value = meta.license || '';
        licenseInput.addEventListener('input', () => {
            item.metadata = item.metadata || {};
            item.metadata.license = licenseInput.value;
        });
        metaLicense.appendChild(licenseInput);
        propertiesPanel.appendChild(metaLicense);

    }

    function createRow(labelText, id) {
        const row = document.createElement('div');
        row.className = 'prop-row';
        const label = document.createElement('label');
        label.textContent = labelText;
        label.htmlFor = id;
        label.className = 'prop-label';
        row.appendChild(label);
        return row;
    }

    // Save/export functionality
    const topButtons = document.querySelectorAll('.imagepack-topbar .imagepack-btn');
    const actualSaveBtn = topButtons[2];
    if (actualSaveBtn) {
        actualSaveBtn.addEventListener('click', () => {
            const blob = new Blob([JSON.stringify(pack, null, 2)], {type: 'application/json'});
            const url = URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = (pack.id || 'imagepack') + '.json';
            document.body.appendChild(a);
            a.click();
            a.remove();
            URL.revokeObjectURL(url);
        });
    }

    // Initial render
    renderPreview();
    // Auto-select first image so inputs are visible immediately
    if (pack.content && pack.content.length > 0) selectImage(0);
});
