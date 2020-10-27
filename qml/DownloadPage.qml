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
        title: 'Download ' + u.getCurrentNodeName()
        leadingActionBar {
            actions: [ backAction ]
        }
    }

    Timer {
        id: timer
        interval: 100; repeat: true
        running: true
        triggeredOnStart: true
        onTriggered: {
            console.log(u.percent)
            progressBar.value = u.percent
        }
    }

    ProgressBar {
        id: progressBar
        anchors {
            margins: units.gu(2)
            top: header.bottom
            left: parent.left
            right: parent.right
            //bottom: parent.bottom
        }
        minimumValue: 0
        maximumValue: 100
    }

    Button {
        anchors {
            margins: units.gu(2)
            top: progressBar.bottom
            left: parent.left
            right: parent.right
            //bottom: parent.bottom
        }
        text: 'Download'
        onClicked: {
            u.downloadCurrentNode()
        }
    }
}
