// template: automation
// desc: Basic system automation script (executing shell commands, file io)

pt "--- System Automation Script ---"

 // Get current user from env
user = env("USERNAME")
if len(user) = 0 {
    user = env("USER") // Linux/Mac fallback
}
pt "Running as: ${user}"

 // Create a backup folder
backup_dir = "./backup_" + str(now())
pt "Creating backup directory: ${backup_dir}"
cmd("mkdir", backup_dir)

 // Write a log file
log_file = backup_dir + "/run.log"
w(log_file, "Backup started at " + date()["format"])

 // List current directory
files = dir(".")
pt "Found ${len(files)} files/folders."

 // Process files
lp len(files) {
    f = files[i]
    if has(f, ".txt") {
        pt "Found text file: ${f}"
         // In a real script, you might copy it using cmd("cp", f, backup_dir)
    }
}

pt "Automation complete!"
