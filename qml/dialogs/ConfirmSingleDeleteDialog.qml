import QtQuick 2.4
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3

import "templates"

ConfirmDialog {
    id: rootItem

    destructiveDialog: true

    title: i18n.tr("Delete")
    text: i18n.tr("Are you sure you want to permanently delete '%1'?").arg(u.getCurrentNodeName())
}
