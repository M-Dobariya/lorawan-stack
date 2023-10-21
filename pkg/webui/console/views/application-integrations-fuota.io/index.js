// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Col, Container, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useParams } from 'react-router-dom'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Link from '@ttn-lw/components/link'
import PageTitle from '@ttn-lw/components/page-title'
import Message from '@ttn-lw/lib/components/message'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import React, { useCallback, useState } from 'react'
import FuotaIOImage from '@assets/misc/fuotaio.svg'
import style from './application-integrations-fuota.io.styl'
import FuotaIoForm from '@console/containers/fuota-io-form'
import { useDispatch } from 'react-redux'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import { createApplicationApiKey } from '@console/store/actions/api-keys'
import Button from '@ttn-lw/components/button'
import DataSheet from '@ttn-lw/components/data-sheet'
import Collapse from '@ttn-lw/components/collapse'

const m = defineMessages({
  fuotaIoInfoText:
    'Simplify IoT device management with our cutting-edge IT product. Integrate devices, standardize types, add preferred LNS. Hassle-free firmware updates keep devices optimized. Empower your IoT ecosystem effortlessly and stay up to date.',
  officialFuotaIoWebsite: 'Official FUOTA.IO website',
  fuotaIoApiKeys: 'FUOTA.IO API Keys',
  tokenDescription: 'Fuota access token as configured within FUOTA.IO',
  generateApiKey: 'Generate new API key',
  goToApiKeys: 'Go to API keys',
  generateFuotaIoDescription:
    'Generate a new API key to connect your FUOTA.IO account to The Things Stack.',
  setFuotaIoDescription:
    'Set FUOTA.IO API Key to connect your FUOTA.IO account to The Things Stack.',
})

const ApplicationIntegrationsFuotaIo = () => {
  const { appId } = useParams()

  const [password, setPassword] = useState('')

  const dispatch = useDispatch()

  useBreadcrumbs(
    'apps.single.integrations.fuotaio',
    <Breadcrumb
      path={`/applications/${appId}/integrations/fuota.io`}
      content={sharedMessages.fuotaIo}
    />,
  )

  const handleGeneratePasswordClick = useCallback(async () => {
    const key = {
      name: `mqtt-password-key-${Date.now()}`,
      rights: ['RIGHT_APPLICATION_TRAFFIC_READ', 'RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE'],
    }
    const result = await dispatch(attachPromise(createApplicationApiKey(appId, key)))
    setPassword(result)
  }, [appId, dispatch])

  const connectionData = [{ header: m.connectionCredentials, items: [] }]

  if (password) {
    connectionData[0].items.push({
      label: m.tokenDescription,
      value: password.key,
      // key: 'Password',
      type: 'code',
    })
  } else {
    connectionData[0].items.push({
      // key: sharedMessages.password,
      value: (
        <>
          <Button
            message={m.generateApiKey}
            onClick={handleGeneratePasswordClick}
            className="mr-cs-s"
          />
          <Link to={`/applications/${appId}/api-keys`} naked secondary>
            <Message content={m.goToApiKeys} />
          </Link>
        </>
      ),
    })
  }

  return (
    <Container>
      <PageTitle title={sharedMessages.fuotaIo} />
      <Row>
        <Col lg={8} md={12}>
          <img className={style.logo} src={FuotaIOImage} alt="LoRa Cloud" />
          <Message content={m.fuotaIoInfoText} className="mt-0" />
          <div>
            <Message
              component="h4"
              content={sharedMessages.furtherResources}
              className="mb-cs-xxs"
            />
            <Link.Anchor href="https://fuota.io/" external secondary>
              <Message content={m.officialFuotaIoWebsite} />
            </Link.Anchor>
          </div>
          {/* <hr className="mb-ls-s" /> */}
          <Collapse
            title={'Generate FUOTA.IO Webhook API key'}
            description={m.generateFuotaIoDescription}
          >
            <DataSheet data={connectionData} />
          </Collapse>
          <Collapse title={'Set FUOTA.IO API key'} description={m.setFuotaIoDescription}>
            <FuotaIoForm />
          </Collapse>
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationIntegrationsFuotaIo
