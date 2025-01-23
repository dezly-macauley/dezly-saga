<script lang="ts">
  let text_with_html_tags: string = "The word <strong>ninja</strong> is in bold.";
</script>

<p>{@html text_with_html_tags}</p>

<!-- 
  NOTE: Without the `@html` Svlete will display the html tag instead of 
  applying the effect:
  The word <strong>ninja<strong> is in bold.
-->

<!-- <p>{text_with_html_tags}</p> -->

<!-- 
  WARNING: 
 Important: Svelte doesn’t perform any sanitization of the expression 
 inside {@html ...} before it gets inserted into the DOM.
 This isn’t an issue if the content is something you trust 
 like an article you wrote yourself.
 However if it’s some untrusted user content,
 e.g. a comment on an article,
 then it’s critical that you manually escape it,
 otherwise you risk exposing your users 
 to Cross-Site Scripting (XSS) attacks.

-->
