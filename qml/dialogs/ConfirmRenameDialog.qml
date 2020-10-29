import QtQuick 2.4
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3

import "templates"

ConfirmDialogWithInput {
    id: rootItem

    title: i18n.tr("Rename")
    text: i18n.tr("Enter a new name")
}
