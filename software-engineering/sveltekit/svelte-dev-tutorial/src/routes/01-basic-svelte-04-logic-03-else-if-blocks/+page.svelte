<script lang="ts">
  let count: number = $state(0);
  
  function increment() {
    count += 1;
  }
</script>

<button onclick={increment}>
  Clicked {count}
  {count == 1? "time" : "times"}
</button>

<!-- 
  NOTE: This button will only appear if count is greater than 10 
-->
{#if count > 10}
  <p>{count} is greater than 10</p>
{:else if count == 10}
  <p>{count} is exactly 10</p>
{:else}
  <p>{count} is less than 10</p>
{/if}
