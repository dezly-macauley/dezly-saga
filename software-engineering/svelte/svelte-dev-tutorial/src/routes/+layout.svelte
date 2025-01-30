<script lang="ts">

  // NOTE: Uncomment this to make sute that every 
  // child route (+pages.svelte) has the css style

  // import '../app.css';

  let { children } = $props();
</script>

<h1>Svelte Dev Tutorial</h1>
<nav>
  <a href="/">Home</a>
</nav>

<hr>

<!-- 
  NOTE: This is a convinient way of making sure
  that every child route (+page.svelte) has the heading above.
-->
{@render children()}
