# ABOUT: My SvelteKit Setup
# -----------------------------------------------------------------------------
# To make this script executable, run this command first:
# chmod +x bun_vite_html_css_typescript_sveltekit_setup.sh
# -----------------------------------------------------------------------------
# Run the script with this command:
# source bun_vite_html_css_typescript_sveltekit_setup
# -----------------------------------------------------------------------------

#______________________________________________________________________________

project_name=""

echo "Name of new project:"
read project_name 

bunx sv create --check-types typescript --template minimal $project_name

cd $project_name

echo "To start and open the server in the browser, run this command:"
echo "bun run dev --open"

#______________________________________________________________________________
