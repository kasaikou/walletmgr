import sys

import gi

gi.require_version("Gtk", "4.0")
from gi.repository import GLib, Gtk

class WalletManagerGUIApplication(Gtk.Application):
    def __init__(self):
        super().__init__(application_id="com.kasaikou.walletmgr")
        GLib.set_application_name("walletmgr")
    
    def do_activate(self):
        window = Gtk.ApplicationWindow(application=self, title="Wallet Manager")
        window.present()
        
app = WalletManagerGUIApplication()
sys.exit(app.run(sys.argv))
