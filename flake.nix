{
  description = "A link redirect server written in Go";

  inputs = {
    nixpkgs.url = "github:nixOS/nixpkgs/release-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
  	flake-utils.lib.eachDefaultSystem (system:
			let
				name = "redirgo";
				src = ./.;
				pkgs = import nixpkgs {
					inherit system;
				};
			in with pkgs; {
				devShells.default = mkShell {
					inherit name;
					buildInputs = with pkgs; [
						go
					];
					shellHook = ''
						echo "Entered development environment"
					'';
				};
			}
		);
}
