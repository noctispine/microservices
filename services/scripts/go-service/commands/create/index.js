const fs = require("fs")
const path = require("path")
const inquirer = require('inquirer')
const chalk = require("chalk")
const { getAllFiles, constructFileDirs } = require("./helper")
const { questions } = require('./constants')

async function create(program) {
    prompt = inquirer.createPromptModule()
    let answers

    await prompt(questions).then(inputs => {
        answers = { ...inputs }
    })

    const { projectName, port } = answers
    
    copyProject(projectName)
    const files = getAllFiles(`./${projectName}`)
    files.push(`./proto-buffers/${projectName}.proto`)
    
    console.log(files)

    files.forEach(file => {
        fs.readFile(file, 'utf-8', (err, data) => {
            if(err) {
                console.error("error occured while reading a file", err)
                process.exit(1)
            }

            const result = data
                .replace(/\$serviceNameCapitalized/g, `${projectName.charAt(0).toUpperCase() + projectName.slice(1)}`)
                .replace(/\$serviceName/g, projectName)
                .replace(/\$PORT/g, port)
            
            fs.writeFile(file, result, 'utf-8', err => {
                if(err) {
                    console.error("error occured while writing to file", err)
                    process.exit(1)
                }
            })

        })
    })

    console.log(chalk.yellow(projectName) + chalk.cyan(" service is created"))
}

function copyProject(projectName) {
    const fileDirs = constructFileDirs(projectName)
    fileDirs.forEach(fileDir => {
        fs.cpSync(fileDir.src, fileDir.dest, {recursive: true}, (err) => {
            console.error(`${path.basename(fileDir.src)} cannot copied to destination`, err)
            process.exitCode(2)
        })
    })
}

module.exports = { create }