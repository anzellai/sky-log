"""Help must finish before the web server starts or configuration is read."""
from pathlib import Path
import subprocess
import tempfile

binary = Path(__file__).resolve().parents[1] / "sky-out" / "app"
with tempfile.TemporaryDirectory(prefix="sky-log-help-") as directory:
    for argument in ["--help", "-h"]:
        result = subprocess.run(
            [str(binary), argument], cwd=directory,
            env={"HOME": directory, "PATH": "/usr/bin:/bin"},
            capture_output=True, text=True, timeout=5,
        )
        assert result.returncode == 0, "help command did not exit successfully"
        assert "Usage:" in result.stdout, "help text was not emitted"
        assert "listening" not in result.stdout.lower(), "help started a server"
print("Help startup regression: PASS")
