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
import React from 'react'
import FuotaIOImage from '@assets/misc/fuotaio.svg'
import style from './application-integrations-fuota.io.styl'
import DataSheet from '@ttn-lw/components/data-sheet'

const m = defineMessages({
  fuotaIoInfoText:
    'Simplify IoT device management with our cutting-edge IT product. Integrate devices, standardize types, add preferred LNS. Hassle-free firmware updates keep devices optimized. Empower your IoT ecosystem effortlessly and stay up to date.',
  officialFuotaIoWebsite: 'Official FUOTA.IO website',
  fuotaIoApiKeys: 'FUOTA.IO API Keys',
})

const ApplicationIntegrationsFuotaIo = () => {
  const { appId } = useParams()

  useBreadcrumbs(
    'apps.single.integrations.fuotaio',
    <Breadcrumb
      path={`/applications/${appId}/integrations/fuota.io`}
      content={sharedMessages.fuotaIo}
    />,
  )

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
          <hr className="mb-ls-s" />
          <Message content={m.fuotaIoApiKeys} component="h3" />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationIntegrationsFuotaIo
