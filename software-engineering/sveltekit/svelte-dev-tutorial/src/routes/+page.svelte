<h2>Basic Svelte</h2>

<nav>

  <a href="/01-basic-svelte-01-introduction-01-welcome-to-svelte">
  01 Introduction, 01 Welcome to Svelte
  </a>

  <a href="/01-basic-svelte-01-introduction-02-your-first-component">
  01 Introduction, 02 Your First Component
  </a>

  <a href="/01-basic-svelte-01-introduction-03-dynamic-attributes">
  01 Introduction, 03 Dynamic attributes
  </a>

  <a href="/01-basic-svelte-01-introduction-04-styling">
  01 Introduction, 04 Styling
  </a>

</nav>

<hr>

<h2>Basic SvelteKit</h2>

<nav>

  <a href="/03-basic-sveltekit-02-routing-01-pages">
  02 Routing, 01 Pages
  </a>

</nav>

<hr>

<!---------------------------------------------------------------------------->
<!-- SECTION: CSS -->

<style>
  nav {

    /*  NOTE: Flexbox

      1. This will enable the flexbox layout,
      and turn the `nav` container into a flex container

      2. The child elements in the container (the <a> tags,
      become flex items.

      3. You can then use flex-direction to tell the browser how to
      display the items inside the container.

      4. flex-direction: columns;` means "display the anchor tags
      in a column (On top of each other).

    */
    display: flex;
    flex-direction: column;
  }
</style>

<!---------------------------------------------------------------------------->
