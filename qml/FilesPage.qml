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
//import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0

Page {
    id: filesPage
    anchors.fill: parent

    property Action backAction: Action {
        id: backAction
        objectName: "backButton"
        iconName: "back"
        onTriggered: {
            console.log("Exit from folder")
            u.popNode()
            u.regenerateDictionary()
            pageStack.pop()
        }
    }

    header: PageHeader {
        id: header
        title: i18n.tr('MEGAClient')
        leadingActionBar {
            actions: {
                if (pageStack.depth == 1)
                    return null
                else
                    return [ backAction ]
            }
        }
    }

    ListView {
        id: filesListView
        spacing: units.gu(2)
        anchors {
            margins: units.gu(2)
            top: header.bottom
            left: parent.left
            right: parent.right
            bottom: parent.bottom
        }

        delegate: Item {
            id: item
            anchors.left: parent.left
            anchors.right: parent.right
            height: 40
            property string buttonHash: hash
            Button {
                anchors.fill: parent
                anchors.margins: 5
                text: buttonText
                onClicked: {
                    u.pushNode(buttonHash)
                    if(!u.getNumberOfChildren())
                        PopupUtils.open(Qt.resolvedUrl("dialogs/ActionsDialog.qml"), mainView, {"currentPage": filesPage})
                    else
                        pageStack.push(Qt.resolvedUrl("FilesPage.qml"))
                }
            }
        }

        model: ListModel {
            id: listModel // задаём ей id для обращения
        }
    }

    function makeButtons() {
        listModel.clear();
        var paths_as_str = u.getFiles();
        var paths = paths_as_str.split("|")
        var hashes_as_str = u.getHashes()
        var hashes = hashes_as_str.split("|")
        if(paths.length == 0)
            console.log("Empty paths");
        else
            paths.forEach(function(item, i, paths) {
                listModel.append({buttonText: item, hash: hashes[i]})
            });
    }

    Component.onCompleted: {
        makeButtons()
        header.title = u.getCurrentNodeName();
    }
}
