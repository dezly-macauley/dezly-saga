<script>
    
    // NOTE: The `static` directory

    // In SvelteKit when you want to reference a static file,
    // you need to use the root-relative path starting 
    // from the static directory
    let vegeta_image = "/vegeta.gif";

</script>

<img 
    src={vegeta_image} 
    alt="Super saiyain vegeta"
    height="200"
    width="300"
>


