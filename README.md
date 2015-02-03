# GoBuilderIDEA
Simple external tool for IDEA+Go language plugin. Just run "go install" for selected file

# How to install
1. Build util (simple use go install in GoBuilderIDEA dir)
- Go to Settings->Tools->External Tools
- Add new tool
- Set name: GoBuilder
- Set Program:  change util bin file
- Set Parameters: $ProjectFileDir$
- Set Working directory: $FileDir$

# How use
1. Select the file you want to install
- In right menu select External Tools->GoBuilder
