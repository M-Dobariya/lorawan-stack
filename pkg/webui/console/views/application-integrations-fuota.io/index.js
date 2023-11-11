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

import FuotaIOImage from '@assets/misc/fuotaio.svg'
import FuotaIoForm from '@console/containers/fuota-io-form'
import Require from '@console/lib/components/require'
import { mayViewFuotaIoInfo } from '@console/lib/feature-checks'
import { createApplicationApiKey } from '@console/store/actions/api-keys'
import SubViewError from '@console/views/sub-view-error'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Button from '@ttn-lw/components/button'
import Collapse from '@ttn-lw/components/collapse'
import DataSheet from '@ttn-lw/components/data-sheet'
import Link from '@ttn-lw/components/link'
import PageTitle from '@ttn-lw/components/page-title'
import { ErrorView } from '@ttn-lw/lib/components/error-view'
import Message from '@ttn-lw/lib/components/message'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import React, { useCallback, useState } from 'react'
import { Col, Container, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useDispatch } from 'react-redux'
import { useParams } from 'react-router-dom'
import style from './application-integrations-fuota.io.styl'
import RequireRequest from '@ttn-lw/lib/components/require-request'
import { getAppPkgAssoc } from '@console/store/actions/application-packages'

const m = defineMessages({
  goToApiKeys: 'Go to API keys',
  fuotaIoApiKeys: 'FUOTA.IO API Keys',
  generateApiKey: 'Generate new API key',
  officialFuotaIoWebsite: 'Official FUOTA.IO website',
  tokenDescription: 'Fuota access token as configured within FUOTA.IO',
  setFuotaIoDescription:
    'Set FUOTA.IO API Key to connect your FUOTA.IO account to The Things Stack.',
  generateFuotaIoDescription:
    'Generate a new API key to connect your FUOTA.IO account to The Things Stack.',
  fuotaIoInfoText:
    'Simplify IoT device management with our cutting-edge IT product. Integrate devices, standardize types, add preferred LNS. Hassle-free firmware updates keep devices optimized. Empower your IoT ecosystem effortlessly and stay up to date.',
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
      name: `fuota-io-api-key-${Date.now()}`,
      rights: ['RIGHT_APPLICATION_TRAFFIC_READ', 'RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE'],
    }
    const result = await dispatch(attachPromise(createApplicationApiKey(appId, key)))
    setPassword(result)
  }, [appId, dispatch])

  const connectionData = [{ header: m.connectionCredentials, items: [] }]

  const selector = ['data']

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
    <Require featureCheck={mayViewFuotaIoInfo} otherwise={{ redirect: `/applications/${appId}` }}>
      <RequireRequest
        requestAction={[
          getAppPkgAssoc(appId, 207, 1, selector),
          // getAppPkgAssoc(appId, LORA_CLOUD_GLS.DEFAULT_PORT, selector),
        ]}
      >
        <ErrorView errorRender={SubViewError}>
          <Container>
            <PageTitle title={sharedMessages.fuotaIo} />
            <Row>
              <Col lg={8} md={12}>
                <img className={style.logo} src={FuotaIOImage} alt="fuota.io" />
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
        </ErrorView>
      </RequireRequest>
    </Require>
  )
}

export default ApplicationIntegrationsFuotaIo
