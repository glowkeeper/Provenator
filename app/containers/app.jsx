import React from 'react'
import {Route, Link} from 'react-router-dom'
import { Icon, Layout, Menu, Breadcrumb, Row, Col } from 'antd'

import {logo} from '../images'

import Web3Handler from '../utils/web3Handler'
import ContractHandler from '../utils/contractHandler'
import {AppStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings} from '../helpers/outputStrings'
import {AppURLS} from '../helpers/urls'
import {AppPaths} from '../helpers/paths'
import {AppHeading, IOTagline} from '../components/io'

import Events from './helpers/events'
import Home from './helpers/home'
import About from './helpers/about'
import Overview from './helpers/overview'
import Help from './helpers/help'
import Writer from './writer/writer'
import Reader from './reader/reader'

class App extends React.Component {

  constructor (props) {
    super(props)

    this.web3Handler = new Web3Handler()
    this.contractHandler = new ContractHandler(this.web3Handler)
  }

  render () {

    const { SubMenu } = Menu
    const { Header, Content, Sider, Footer } = Layout

    return (
      <div>
        <Layout>
          <Header>
            <Row>
              <Col span={4}><img src={logo} /></Col>
              <Col span={10}>
                <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }} >
                  <Menu.Item key={AppStrings.home}>
                    <Icon type={HomeStrings.icon}/><span>{AppStrings.home}</span>
                    <Link to={AppPaths.home}/>
                  </Menu.Item>
                  <Menu.Item key={AppStrings.about}>
                    <Icon type={AboutStrings.icon}/><span>{AppStrings.about}</span>
                    <Link to={AppPaths.about}/>
                  </Menu.Item>
                  <Menu.Item key={AppStrings.overview}>
                    <Icon type={OverviewStrings.icon}/><span>{AppStrings.overview}</span>
                    <Link to={AppPaths.overview}/>
                  </Menu.Item>
                  <Menu.Item key={AppStrings.help}>
                    <Icon type={HelpStrings.icon}/><span>{AppStrings.help}</span>
                    <Link to={AppPaths.help}/>
                  </Menu.Item>
                </Menu>
              </Col>
              <Col span={10} style={{ textAlign: 'right' }}>
                <h5>{AppStrings.tagLine}</h5>
              </Col>
            </Row>
          </Header>
        </Layout>
        <Layout>

          <Sider width={200}>
            <Menu mode="inline" defaultSelectedKeys={['1']} defaultOpenKeys={['Home']} style={{ height: '100%', borderRight: 0 }} >
              <Menu.Item key={AppStrings.create}>
                <Icon type={AppStrings.createIcon}/><span>{AppStrings.create}</span>
                <Link to={AppPaths.create}/>
              </Menu.Item>
              <Menu.Item key={AppStrings.read}>
                <Icon type={AppStrings.readIcon}/><span>{AppStrings.read}</span>
                <Link to={AppPaths.read}/>
              </Menu.Item>
            </Menu>
          </Sider>

          <Layout style={{ padding: '0 24px 24px' }}>

            <Breadcrumb style={{ margin: '16px 0' }}>
              <Breadcrumb.Item>{AppStrings.home}</Breadcrumb.Item>
              <Breadcrumb.Item>{AppStrings.create}</Breadcrumb.Item>
              <Breadcrumb.Item>{AppStrings.read}</Breadcrumb.Item>
              <Breadcrumb.Item>{AppStrings.about}</Breadcrumb.Item>
              <Breadcrumb.Item>{AppStrings.overview}</Breadcrumb.Item>
              <Breadcrumb.Item>{AppStrings.help}</Breadcrumb.Item>
            </Breadcrumb>

            <Content style={{ padding: 24, margin: 0, minHeight: 280 }}>

              <Route name={AppStrings.home} exact path={AppPaths.home} component={Home} />
              <Route name={AppStrings.about} path={AppPaths.about} component={About} />
              <Route name={AppStrings.overview} path={AppPaths.overview} component={Overview} />
              <Route name={AppStrings.help} path={AppPaths.help} component={Help} />
              <Route
                name={AppStrings.create}
                path={AppPaths.create}
                render={() => <Writer contracts={this.contractHandler} web3={this.web3Handler} />}
              />
              <Route
                name={AppStrings.read}
                path={AppPaths.read}
                render={() => <Reader contracts={this.contractHandler} web3={this.web3Handler} />}
              />

            </Content>
          </Layout>
        </Layout>
        <Layout>
          <Footer style={{ textAlign: 'center' }}>
            <h5>{AppStrings.copyright}</h5>
          </Footer>
        </Layout>
      </div>

    )
  }
}

/* <li><Link to="/events">FakeNews Events</Link></li>
<Route path="/events" render={() => <Events contracts={this.contractHandler} web3={this.web3Handler} />} /> */

/*
<div>
  <Layout>
    <Header>
      <Row>
        <Col span={11}>

          <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }} >
            <Menu.Item key={AppStrings.home}>
              <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.home}</span>
              <Link to={AppPaths.home}/>
            </Menu.Item>
            <Menu.Item key={AppStrings.about}>
              <Icon type={AboutStrings.aboutIcon}/><span>{AppStrings.about}</span>
              <Link to={AppPaths.about}/>
            </Menu.Item>
            <Menu.Item key={AppStrings.overview}>
              <Icon type={OverviewStrings.overviewIcon}/><span>{AppStrings.overview}</span>
              <Link to={AppPaths.overview}/>
            </Menu.Item>
            <Menu.Item key={AppStrings.help}>
              <Icon type={HelpStrings.helpIcon}/><span>{AppStrings.help}</span>
              <Link to={AppPaths.help}/>
            </Menu.Item>
          </Menu>
        </Col>
          <Col span={11} style={{ textAlign: 'right' }}>{AppStrings.tagLine}</Col>
      </Row>
    </Header>
    <Layout>
      <Sider width={200}>
        <Menu mode="inline" defaultSelectedKeys={['1']} defaultOpenKeys={['Home']} style={{ height: '100%', borderRight: 0 }} >
          <Menu.Item key={HomeStrings.home}>
            <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.home}</span>
            <Link to={AppPaths.home}/>
          </Menu.Item>

          <SubMenu title={<span><Icon type={HomeStrings.homeIcon}/><span>{AppStrings.create}</span></span>}>
              <Menu.Item key={AppStrings.create}>
                <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.create}</span>
                <Link to={AppPaths.create}/>
            </Menu.Item>
          </SubMenu>

          <SubMenu title={<span><Icon type={HomeStrings.homeIcon} /><span>{AppStrings.read}</span></span>}>
            <Menu.Item key={AppStrings.read}>
                <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.read}</span>
                <Link to={AppPaths.read}/>
              </Menu.Item>
          </SubMenu>

          <Menu.Item key={AppStrings.about}>
            <Icon type={AboutStrings.aboutIcon}/><span>{AppStrings.about}</span>
            <Link to={AppPaths.about}/>
          </Menu.Item>'../helpers/outputStrings'
          <Menu.Item key={AppStrings.overview}>
            <Icon type={OverviewStrings.overviewIcon}/><span>{AppStrings.overview}</span>
            <Link to={AppPaths.overview}/>
          </Menu.Item>
          <Menu.Item key={AppStrings.help}>
            <Icon type={HelpStrings.helpIcon}/><span>{AppStrings.help}</span>
            <Link to={AppPaths.help}/>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout style={{ padding: '0 24px 24px' }}>
        <Breadcrumb style={{ margin: '16px 0' }}>
          <Breadcrumb.Item>Home</Breadcrumb.Item>
          <Breadcrumb.Item>Create</Breadcrumb.Item>
          <Breadcrumb.Item>Read</Breadcrumb.Item>
          <Breadcrumb.Item>About</Breadcrumb.Item>
          <Breadcrumb.Item>Overview</Breadcrumb.Item>
          <Breadcrumb.Item>Help</Breadcrumb.Item>
        </Breadcrumb>
        <Content style={{ padding: 24, margin: 0, minHeight: 280 }}>

          <Route name={AppStrings.home} exact path={AppPaths.home} component={Home} />
          <Route name={AppStrings.about} path={AppPaths.about} component={About} />
          <Route name={AppStrings.overview} path={AppPaths.overview} component={Overview} />
          <Route name={AppStrings.help} path={AppPaths.help} component={Help} />

          <Route
            name={AppStrings.create}
            path={AppPaths.create}
            render={() => <Writer contracts={this.contractHandler} web3={this.web3Handler} />}
          />
          <Route
            name={AppStrings.read}
            path={AppPaths.read}
            render={() => <Reader contracts={this.contractHandler} web3={this.web3Handler} />}
          />

        </Content>
        <Footer style={{ textAlign: 'center' }}>
          {AppStrings.copyright}
        </Footer>
      </Layout>
    </Layout>
  </Layout>
</div>
*/

export default App
