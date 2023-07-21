import pynvim
from pynvim import Nvim
import asyncio

from typing import List


@pynvim.plugin
class DenopsIssue268(object):
    def __init__(self, nvim: Nvim):
        self.nvim = nvim

    @pynvim.command("DenopsIssue268Reproduce", sync=False)
    def reproduce(self):
        asyncio.gather(
            self.echo("1"),
            self.echo("2"),
            self.echo("3"),
            self.echo("4"),
            self.echo("5"),
            self.echo("6"),
            self.echo("7"),
            self.echo("8"),
                # asyncio.to_thread(self.echo, "1"),
                # asyncio.to_thread(self.echo, "2"),
                # asyncio.to_thread(self.echo, "3"),
                # asyncio.to_thread(self.echo, "4"),
                # asyncio.to_thread(self.echo, "5"),
                # asyncio.to_thread(self.echo, "6"),
                # asyncio.to_thread(self.echo, "7"),
                # asyncio.to_thread(self.echo, "8"),
        )

    async def echo(self, msg):
        self.nvim.call("DenopsIssue268Echo", msg)

    @pynvim.function("DenopsIssue268Echo", sync=False)
    def echoFunction(self, args: List[str]):
        self.nvim.command(f"echomsg 'Hello, {args[0]}!'")
