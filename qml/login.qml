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

MainView {
    id: loginView
    objectName: 'loginView'
    applicationName: 'mega.roker2'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    Page {
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

            TextInput {
                id: login
                text: "Login"
                color: 'red'
                Layout.alignment: Qt.AlignCenter
                onEditingFinished: { u.login = login.text }
            }

            TextInput {
                id: pswd
                text: "Password"
                color: 'red'
                Layout.alignment: Qt.AlignCenter
                echoMode: TextInput.PasswordEchoOnEdit
                onEditingFinished: { u.password = pswd.text }
            }

            Label {
                text: u.output
                horizontalAlignment: Label.AlignHCenter
                Layout.fillWidth: true
            }

            Button {
                Text {
                    text: "Sign in"
                    color: "white"
                    anchors.centerIn: parent
                }
                color: 'red'
                onClicked: u.signIn()
                Layout.alignment: Qt.AlignCenter
            }
        }
    }
}
