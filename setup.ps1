# Define the project directory
$projectDir = "C:\Users\jluker\VSCodeProjects\verden"

# Create frontend directory
New-Item -Path "$projectDir/frontend" -ItemType Directory

# Initialize the Svelte project in the frontend directory
Set-Location -Path "$projectDir/frontend"
npx degit sveltejs/template svelte-app
Set-Location -Path "svelte-app"
npm install

# Add .gitignore for the frontend
"node_modules/" | Out-File -FilePath "$projectDir/frontend/.gitignore"

# Add .gitignore for the backend (Go)
"bin/
obj/
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
" | Out-File -FilePath "$projectDir/app/.gitignore"

Write-Host "Project setup complete."
