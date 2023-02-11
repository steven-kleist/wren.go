foreign class Platform {
  foreign static name
  static isPosix { name != "windows" }
  static isWindows { name == "windows" }
}
foreign class Process {
  foreign static allArguments
  static arguments { allArguments.count >= 2 ? allArguments[2..-1] : [] }
  foreign static cwd
  foreign static exit_(exit_code)
  foreign static exec(cmdline)
  static exit() {
    exit_(0)
  }
  static exit(code) {
    exit_(code)
  }
}