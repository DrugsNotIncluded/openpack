# openpack
openpack is a command line tool for creating Minecraft modpacks. Instead of managing JAR files directly, openpack creates TOML metadata files which can be easily version-controlled and shared with git. You can then export it to a CurseForge or Modrinth modpack, or use openpack-installer for an auto-updating MultiMC instance.

# TODO:
1. Add subcommand: `automodpack genconfig` generates an [https://github.com/Skidamek/AutoModpack](AutoModpack) configuration with server-only mods excluded and mods marked as server-side on which client-side mods depend.
2. Port export to PrismLauncher format from packwiz-install
3. Write simple GUI in imgui
