<script lang="ts">
  // This is a prop that has a default value.
  // If it is used without setting the player name then the player name will
  // be set to "Mystery Fox" 
  let {player_name = "Mystery Fox"} = $props();
</script>

<!--
  NOTE: Essentially this is a resuable component that can be used
 by suppling it the name of a player.
-->
<p>The player's name is {player_name}</p>
