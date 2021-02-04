{ src ? builtins.fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-20.09.tar.gz",
  pkgs ? import src {}}:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_15
    gopls
    delve
    go-outline

    mysql80.client
    postgresql_10
  ];

  hardeningDisable = [ "all" ];

  GO111MODULE = "on";

  shellHook = ''
    PATH=~/go/bin:$PATH

    if [ ! -d tests/testdata/.git ]; then
      git submodule init
    fi

    git submodule update
  '';
}
