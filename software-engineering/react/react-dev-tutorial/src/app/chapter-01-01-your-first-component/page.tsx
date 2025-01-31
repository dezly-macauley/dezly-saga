//_____________________________________________________________________________
// SECTION: Step 1 - Create the component

// NOTE: `function Profile(){return();}`
// ---------------------------------------
// 1. A React component is a TypeScript / JavaScript function 
// that returns multiple lines of HTML
// 2. The name of the React component must start with a capital letter 
// 3. The html contents must be placed inside `return();` 

// The reason why there is no `export default` before the word 
// function is because I only need this component to be used within this file
function Profile() {
  return (
    <img
      src="/katherine.png" 
      alt="Katherine Johnson"
    />
  );
}

//_____________________________________________________________________________
// SECTION: Step 2 - Create the Amazing


// NOTE: `export default`
// ------------------------------
// 1. This is standard TypeScript / JavaScript syntax (Not specific to React)
// 2. It lets you mark the main function in a file so that you can later
// import it into another file.

export default function Gallery() {
  return (
    <section>
      <h1>Amazing scientists</h1>

      {/* NOTE: This is how to to use a React component
      */}

      <Profile />
      <Profile />
      <Profile />
    </section>
  );
}
