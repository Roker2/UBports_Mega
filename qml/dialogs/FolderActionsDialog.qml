/*
 * Copyright (C) 2020  Dmitry Minko
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * MEGA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import QtQuick 2.7
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3

Dialog {
    id: actionsDialog

    Label {
        text: u.getCurrentNodeName()
    }

    Button {
        id: downloadButton
        text: "Download"
        onClicked: {
            pageStack.push(Qt.resolvedUrl("../DownloadPage.qml"))
            PopupUtils.close(actionsDialog)
        }
    }

    Button {
        id: renameButton
        text: "Rename"
        onClicked: {
            var props = {
                "inputText": u.getCurrentNodeName()
            }
            var popup = PopupUtils.open(Qt.resolvedUrl("ConfirmRenameDialog.qml"), mainView, props)
            popup.accepted.connect(function(inputText) {
                console.log("Rename " + u.getCurrentNodeName() + " to " + inputText)
                u.renameCurrentNode(inputText)
                u.popNode()
                u.regenerateDictionary()
                pageStack.currentPage.makeButtons()
                PopupUtils.close(actionsDialog)
            })
        }
    }

    Button {
        id: deleteButton
        text: "Delete"
        onClicked: {
            var popup = PopupUtils.open(Qt.resolvedUrl("ConfirmSingleDeleteDialog.qml"), mainView)
            popup.accepted.connect(function() {
                console.log("Delete accepted for " + u.getCurrentNodeName())
                u.deleteCurrentNode()
                u.popNode()
                u.regenerateDictionary()
                pageStack.currentPage.makeButtons()
                PopupUtils.close(actionsDialog)
            })
        }
    }

    Button {
        id: cancelButton
        text: "Cancel"
        onClicked: {
            u.popNode()
            PopupUtils.close(actionsDialog)
        }
    }
}