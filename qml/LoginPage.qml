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
    id: loginPage
    signal signIn
    anchors.fill: parent

    header: PageHeader {
        id: header
        title: i18n.tr('MEGAClient')
    }

    ColumnLayout {
        spacing: units.gu(2)
        anchors {
            margins: units.gu(2)
            top: header.bottom
            left: parent.left
            right: parent.right
            bottom: parent.bottom
        }

        Rectangle {
            width: parent.width - 30 ;
            height: 22
            anchors.horizontalCenter: parent.horizontalCenter
            border.width: 3
            border.color: "#000000"
            radius: 5

            TextInput {
                id: login
                x: 5
                y: 2
                width: parent.width - x ;
                //text: "Login"
                onEditingFinished: { u.login = login.text }
            }
        }

        Rectangle {
            width: parent.width - 30 ;
            height: 22
            anchors.horizontalCenter: parent.horizontalCenter
            border.width: 3
            border.color: "#000000"
            radius: 5

            TextInput {
                id: pswd
                x: 5
                y: 2
                width: parent.width - x ;
                //text: "Password"
                echoMode: TextInput.PasswordEchoOnEdit
                onEditingFinished: { u.password = pswd.text }
            }
        }

        Button {
            Text {
                text: "Sign in"
                color: "white"
                anchors.centerIn: parent
            }
            onClicked: loginPage.signIn()
            color: 'red'
            Layout.alignment: Qt.AlignCenter
        }
    }
}
