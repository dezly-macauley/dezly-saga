<script>
  
  // NOTE: This is how you import a Svelte component
  // $lib is an alias for `src/lib`
  // It makes it convinient to import a component without using having to 
  // use a long relative path.
  import Nested from "$lib/components/Nested.svelte"

</script>

  <p>This is a paragraph.</p>

  <!--  NOTE: This is how use a Svelte component -->
  <Nested />

<style>
  /*  
    NOTE: Styles in Svelte are file-scoped
    So this style rules here will only apply to the `p` tags in
    this file. (Imported files will not be affected)
  */
  p {
  color: goldenrod;
    font-family: 'Comic Sans MS', cursive;
    font-size: 2em;
  }

</style>
