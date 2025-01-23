<script lang="ts">

  // NOTE: $state()
  // This is called a rune in Svelte, and it is part of Svelte's system of
  // reactivity.
  // ------------------------------------------------------------------------
  // It tells Svelte that if something updated the value inside the variable
  // click_counter, then Svelte should also update all the parts
  // of the ui (user interface) that depend the variable click_counter.
  // E.g. Like the heading that displays how many times the button has been
  // clicked
  // ------------------------------------------------------------------------
  let click_counter: number = $state(0);

  function increase_click_counter() {
    // Increase the value of click counter by 1
    click_counter += 1;
  }

</script>

<!--
  NOTE: When this button is clicked it will 
  call the function `increase_click_counter`

-->

<h4>
This button has been clicked {click_counter}
<!--
  NOTE: `?` is called a ternary operator
  If click_counter is 1 then the heading will be:
  This button has been clicked 1 time
  ----------------------------------------------------------
  Otherwise it will display:
  This button has been clicked {click_counter} times
  
  //----------------------------------------------------------
-->
{click_counter == 1 ? "time" : "times"}
</h4>
<button onclick={increase_click_counter}>
  Click Me
</button>
