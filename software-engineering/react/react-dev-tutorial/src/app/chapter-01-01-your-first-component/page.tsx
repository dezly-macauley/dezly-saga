// NOTE: `export default`
// 1. This is standard JavaScript syntax (Not specific to React)
// 2. It lets you mark the main function in a file so that you can later
// import it into another file.

// NOTE: `function Profile() {}` 
// This is a standard JavaScript function

// WARNING: React components are regular JavaScript functions but their
// names must start with a capital letter or they won't work.

//_____________________________________________________________________________
// SECTION: A React component that returns one line of markup

// export default function Profile() {
//   // This is a React component that returns one thing
//   return <img src="/katherine.png" alt="Katherine Johnson"/>;
// }

//_____________________________________________________________________________
// SECTION: A React component that returns multiple line of markup
// 1. Use `return();`

function Profile() {
  return (
    <div>
      <img src="/katherine.png" alt="Katherine Johnson"/>
    </div>
  );
}

//_____________________________________________________________________________

export default function Gallery() {
  return(
    <section>
      <h1>Amazing Scientists</h1>
      <Profile />
      <Profile />
      <Profile />
    </section>
  );
}


//_____________________________________________________________________________
