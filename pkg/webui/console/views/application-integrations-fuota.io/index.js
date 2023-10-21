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
