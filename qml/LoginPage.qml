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
import GoIOComponent 0.1

Page {
    id: loginPage
    anchors.fill: parent

    header: PageHeader {
        id: header
        title: i18n.tr('MEGA Client')
    }

    Label {
        id: loginLabel
        anchors {
            margins: units.gu(2)
            left: parent.left
            verticalCenter: inputRecLogin.verticalCenter
        }
        text: "Login:"
    }

    IOComponent {
        id: io
    }

    InputRectangle {
        id: inputRecLogin
        anchors {
            margins: units.gu(2)
            top: header.bottom
            left: loginLabel.right
            right: parent.right
        }
        TextInput {
            id: login
            x: 5
            y: 2
            width: parent.width - x ;
            //text: "Login"
            onEditingFinished: { u.login = login.text }
        }
    }

    Label {
        id: pswdLabel
        anchors {
            margins: units.gu(2)
            left: parent.left
            verticalCenter: inputRecPswd.verticalCenter
        }
        text: "Password:"
    }

    InputRectangle {
        id: inputRecPswd
        anchors {
            margins: units.gu(2)
            top: inputRecLogin.bottom
            left: pswdLabel.right
            right: parent.right
        }
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

    CheckBox {
        id: checkBoxSaveLogin
        anchors {
            margins: units.gu(2)
            top: inputRecPswd.bottom
            left: parent.left
        }
        checked: false
    }

    Label {
        anchors {
            margins: units.gu(2)
            left: checkBoxSaveLogin.right
            verticalCenter: checkBoxSaveLogin.verticalCenter
        }
        text: "Save login"
    }

    Button {
        anchors {
            margins: units.gu(2)
            bottom: parent.bottom
            horizontalCenter: parent.horizontalCenter
        }
        Text {
            text: "Sign in"
            color: "white"
            anchors.centerIn: parent
        }
        onClicked: {
            if (u.signIn()) {
                if (checkBoxSaveLogin.checked)
                    io.writeToFile(0, "login", u.login)
                else
                    io.removeFile(0, "login")
                pageStack.pop()
                pageStack.push(Qt.resolvedUrl("FilesPage.qml"))
            }
        }
        color: 'red'
        Layout.alignment: Qt.AlignCenter
    }

    Component.onCompleted: {
        if (io.fileIsExist(0, "login"))
            login.text = io.readFromFile(0, "login")
    }
}
