<script>
  // Available colours to choose from
  const available_colors = ["red", "orange", "yellow", "green", "blue", "indigo", "violet"];
 
  // This will keep track of the chosen colour
  // Currently set to index 0 (which is red)
  let selected = $state(available_colors[0]);
</script>

<h1 style="color: {selected}">Pick a colour</h1>

<div>

  <!--  
    NOTE: This will create 7 buttons
    One button for each item in available_colors
  -->
  {#each available_colors as color, i}

    <button

    >Click Me</button>

  <!-- TODO: Continue from here  -->

  {/each}

</div>

<button>Click Me</button>

			<!-- style="background: {color}" -->
			<!-- aria-label={color} -->
			<!-- aria-current={selected === color} -->
			<!-- onclick={() => selected = color} -->
		      
		      <!-- {i + 1} -->

<!-- <style> -->
<!-- 	h1 { -->
<!-- 		transition: color 0.2s; -->
<!-- 		font-size: 2rem; -->
<!-- 		font-weight: 700; -->
<!-- 	} -->
<!---->
<!-- 	div { -->
<!-- 		display: grid; -->
<!-- 		grid-template-columns: repeat(7, 1fr); -->
<!-- 		grid-gap: 5px; -->
<!-- 		max-width: 400px; -->
<!-- 	} -->
<!---->
<!-- 	button { -->
<!-- 		aspect-ratio: 1; -->
<!-- 		border-radius: 50%; -->
<!-- 		background: var(--color, #fff); -->
<!-- 		transform: translate(-2px,-2px); -->
<!-- 		filter: drop-shadow(2px 2px 3px rgba(0,0,0,0.2)); -->
<!-- 		transition: all 0.1s; -->
<!-- 		color: black; -->
<!-- 		font-weight: 700; -->
<!-- 		font-size: 2rem; -->
<!-- 	} -->
<!---->
<!-- 	button[aria-current="true"] { -->
<!-- 		transform: none; -->
<!-- 		filter: none; -->
<!-- 		box-shadow: inset 3px 3px 4px rgba(0,0,0,0.2); -->
<!-- 	} -->
<!-- </style> -->
