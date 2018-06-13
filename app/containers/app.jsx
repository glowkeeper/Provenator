import React from 'react'
import {Route, Link} from 'react-router-dom'
import { Icon, Layout, Menu, Breadcrumb, Row, Col } from 'antd'

import Web3Handler from '../utils/web3Handler'
import ContractHandler from '../utils/contractHandler'
import {AppStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings} from '../helpers/outputStrings'
import {AppURLS} from '../helpers/urls'
import {AppPaths} from '../helpers/paths'
import {AppHeading, IOTagline} from '../components/io'

import Home from './home'
import About from './about'
import Overview from './overview'
import Help from './help'
import Writer from './writer'
import Reader from './reader'
import Events from './events'

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

                <SubMenu title={<span><Icon type={HomeStrings.homeIcon}/><span>{AppStrings.accounts}</span></span>}>
                    <Menu.Item key={AppStrings.create}>
                      <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.create}</span>
                      <Link to={AppPaths.create}/>
                  </Menu.Item>
                </SubMenu>

                <SubMenu title={<span><Icon type={HomeStrings.homeIcon} /><span>{AppStrings.assets}</span></span>}>
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
    )
  }
}

/* <li><Link to="/events">FakeNews Events</Link></li>
<Route path="/events" render={() => <Events contracts={this.contractHandler} web3={this.web3Handler} />} /> */

export default App
