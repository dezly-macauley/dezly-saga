<p>This is a paragraph.</p>

<style>

  p {
    color: goldenrod;

    /*  
      This will use `Noto Serif` as the font for the paragraph.
      If `Noto Serif` is not available on the system then 
      it will use any generic `monospace` font
    */
    font-family: "Noto Serif", "monospace";

    /*
      2em means that the font-size will be twice the size 
      of the parent element.

      In this example since the `p` tag is not inside any specific element,
      then that means that the `body` tag is the parent element.

      The font size of the body in most browsers is 16px

      So that means the font size of the paragraph will be `2 x 16px = 32px`
    */
    font-size: 2em;

  }

</style>
