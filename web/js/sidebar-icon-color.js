// sidebar-icon-color.js
// Assigns a unique, consistent color to each sidebar icon and its underline based on its order in the sidebar
// Uses HSL for visually balanced colors, consistent across launches

(function() {
    const sidebarIcons = document.querySelectorAll('.sidebar-icon');
    const baseHue = 210; // starting hue (blue-ish)
    const hueStep = 37;  // step between icons (visually distinct)
    const saturation = 68;
    const lightness = 62;

    sidebarIcons.forEach((icon, i) => {
        let hue = (baseHue + i * hueStep) % 360;
        let color = `hsl(${hue}, ${saturation}%, ${lightness}%)`;
        icon.style.color = color;
        // Set underline color for active/current tab
        let link = icon.closest('a');
        if (link && (link.classList.contains('active') || link.getAttribute('aria-current') === 'page')) {
            link.style.setProperty('--sidebar-underline', color);
        }
    });
})();
