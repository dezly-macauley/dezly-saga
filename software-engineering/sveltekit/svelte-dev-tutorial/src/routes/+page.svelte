<h2>Basic Svelte</h2>

<hr>

<nav>

  <h3>01 Introduction</h3>

  <a href="/01-basic-svelte-01-introduction-01-welcome-to-svelte">
  01 Welcome to Svelte
  </a>

  <a href="/01-basic-svelte-01-introduction-02-your-first-component">
  02 Your First Component
  </a>

  <a href="/01-basic-svelte-01-introduction-03-dynamic-attributes">
  03 Dynamic attributes
  </a>

  <a href="/01-basic-svelte-01-introduction-04-styling">
  04 Styling
  </a>
  
  <a href="/01-basic-svelte-01-introduction-05-nested-components">
  05 Nested Components
  </a>
  
  <a href="/01-basic-svelte-01-introduction-06-html-tags">
  06 HTML tags
  </a>

</nav>

<hr>

<h3>02 Reactivity</h3>
<nav>

  <a href="/01-basic-svelte-02-reactivity-01-state">
  01 State
  </a>

</nav>

<hr>

<nav>

  <h3>03 Props</h3>

  <a href="/01-basic-svelte-03-props-01-declaring-props">
  01 Declaring props
  </a>
  
  <a href="/01-basic-svelte-03-props-02-default-values">
  02 Default Values
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

  h2 {
    color: darkorchid;
  }

</style>

<!---------------------------------------------------------------------------->
