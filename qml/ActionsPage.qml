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
//import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0

Page {
    anchors.fill: parent

    property Action backAction: Action {
        id: backAction
        objectName: "backButton"
        iconName: "back"
        onTriggered: {
            u.popNode()
            u.regenerateDictionary()
            pageStack.pop()
        }
    }

    header: PageHeader {
        id: header
        title: u.getCurrentNodeName()
        leadingActionBar {
            actions: [ backAction ]
        }
    }

    Label {
        id: sizeLabel
        anchors {
            margins: units.gu(2)
            top: header.bottom
            left: parent.left
            right: parent.right
        }
        text: {
            var suffix = "bytes"
            var size = u.getCurrentNodeSize()
            if (size >= 1024)
            {
                size = size / 1024
                suffix = "KB"
            }
            if (size >= 1024)
            {
                size = size / 1024
                suffix = "MB"
            }
            if (size >= 1024)
            {
                size = size / 1024
                suffix = "GB"
            }
            if (size >= 1024)
            {
                size = size / 1024
                suffix = "TB"
            }
            return "Size " + size.toFixed(2) + " " + suffix
        }
    }

    Button {
        id: downloadButton
        anchors {
            margins: units.gu(2)
            top: sizeLabel.bottom
            left: parent.left
            right: parent.right
        }
        text: "Download"
        onClicked: pageStack.push(Qt.resolvedUrl("DownloadPage.qml"))
    }

    Button {
        id: renameButton
        anchors {
            margins: units.gu(2)
            top: downloadButton.bottom
            left: parent.left
            right: parent.right
        }
        text: "Rename"
    }

    Button {
        id: deleteButton
        anchors {
            margins: units.gu(2)
            top: renameButton.bottom
            left: parent.left
            right: parent.right
        }
        text: "Delete"
    }
}