const fs = require("fs")
const path = require("path")

async function listDirContents(filepath) {
    try {
      const files = await fs.promises.readdir(filepath);
      const detailedFilesPromises = files.map(async (file) => {
        let fileDetails = await fs.promises.lstat(path.resolve(filepath, file));
        const { size, birthtime } = fileDetails;
        return { filename: file, "size(KB)": size, created_at: birthtime };
      });
      const detailedFiles = await Promise.all(detailedFilesPromises);
      console.table(detailedFiles);
    } catch (error) {
      console.error("Error occurred while reading the directory!", error);
    }
}

module.exports = { listDirContents }