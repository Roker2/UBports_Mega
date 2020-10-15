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
    id: mainView
    objectName: 'mainView'
    applicationName: 'mega.roker2'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    LoginPage {
        id: loginPage
        onSignIn: {
            if (u.signIn()) {
                firstPage.show()
                loginPage.hide()
            }
        }
    }

    Page {
        id: firstPage
        anchors.fill: parent

        header: PageHeader {
            id: header2
            title: i18n.tr('MEGAClient')
        }

        Label {
            text: "Test"
            horizontalAlignment: Label.AlignHCenter
            Layout.fillWidth: true
        }

    }
}
