let
  # Import nixpkgs if needed
  pkgs = import <nixpkgs> {};
in
  {
    lib ? pkgs.lib,
    buildGoModule ? pkgs.buildGoModule,
    fetchFromGitHub ? pkgs.fetchFromGitHub,
    installShellFiles ? pkgs.installShellFiles,
    # version and vendorSha256 should be specified by the caller
    version ? "latest",
    vendorSha256,
  }:
    buildGoModule rec {
      pname = "openpack";
      inherit version vendorSha256;

      src = ./..;

      nativeBuildInputs = [
        installShellFiles
      ];

      # Install shell completions
      postInstall = ''
        installShellCompletion --cmd openpack \
          --bash <($out/bin/openpack completion bash) \
          --fish <($out/bin/openpack completion fish) \
          --zsh <($out/bin/openpack completion zsh)
      '';

      meta = with lib; {
        description = "A command line tool for editing and distributing Minecraft modpacks, using a git-friendly TOML format";
        homepage = "https://openpack.infra.link/";
        license = licenses.mit;
        mainProgram = "openpack";
      };
    }
