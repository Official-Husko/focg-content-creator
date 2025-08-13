// Sidebar expand/shrink toggle with state persistence
// This script expects the sidebar to have id="sidebar" and the toggle button to have id="sidebarToggle"
document.addEventListener('DOMContentLoaded', function() {
    const sidebar = document.getElementById('sidebar');
    const toggle = document.getElementById('sidebarToggle');
    if (!sidebar || !toggle) return;
    // Reverse default: start shrunk unless user expanded
    if (localStorage.getItem('sidebarShrunk') === null || localStorage.getItem('sidebarShrunk') === 'true') {
        sidebar.classList.add('shrunk');
        sidebar.classList.remove('expanded');
        toggle.querySelector('.material-icons').textContent = 'chevron_right';
    } else {
        sidebar.classList.add('expanded');
        sidebar.classList.remove('shrunk');
        toggle.querySelector('.material-icons').textContent = 'chevron_left';
    }
    toggle.addEventListener('click', function() {
        const isShrunk = sidebar.classList.toggle('shrunk');
        sidebar.classList.toggle('expanded', !isShrunk);
        // Change icon direction
        const icon = toggle.querySelector('.material-icons');
        icon.textContent = isShrunk ? 'chevron_right' : 'chevron_left';
        // Save state
        localStorage.setItem('sidebarShrunk', isShrunk);
    });
});
