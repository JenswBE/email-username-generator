// Based on https://stackoverflow.com/a/75065536
// Set theme to the user's preferred color scheme
function updateTheme() {
  const isDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const colorMode = isDark ? "dark" : "light";
  document.querySelector("html").setAttribute("data-bs-theme", colorMode);
}

// Set theme on load
updateTheme();

// Update theme when the preferred scheme changes
window
  .matchMedia("(prefers-color-scheme: dark)")
  .addEventListener("change", updateTheme);
