import React from 'react'

import { getKey, getString } from './bytesToString'

import { CreativeWorksProps } from '../store/types'

import { FormHelpers, File, Artefacts } from '../config'

/*
authorId: string
dateCreated: string
dateModified: string
*/

export const getArtefactsHTML = (data: CreativeWorksProps): any[] => {

    let artefactsInfo = []

    //console.log(data)

    if( typeof data.creativeWorks !== 'undefined' ) {
        for ( var i = 0; i < data.creativeWorks.length; i++) {

            const renderArtefactsHTML = (
                <React.Fragment key={data.creativeWorks[i].id}>
                <p>
                    {Artefacts.id}: {data.creativeWorks[i].id}<br/>
                    {Artefacts.artefactName}: {data.creativeWorks[i].name}<br/>
                    {Artefacts.url}: {data.creativeWorks[i].url}<br/>
                    {Artefacts.description}: {data.creativeWorks[i].description}<br/>
                    {Artefacts.dateCreated}: {data.creativeWorks[i].dateCreated}<br/>
                    {Artefacts.dateModified}: {data.creativeWorks[i].dateModified}<br/>
                    {Artefacts.workType}: {FormHelpers.works[data.creativeWorks[i].workType]}<br/>
                    {Artefacts.license}: {FormHelpers.licenses[data.creativeWorks[i].license]}
                </p>
                </React.Fragment>
            )
            artefactsInfo.push(renderArtefactsHTML)

            if( data.creativeWorks[i].author !== null ) {
                if( typeof data.creativeWorks[i].author !== 'undefined' ) {
                    //console.log(data.creativeWorks[i].works)

                    for ( var j = 0; j < data.creativeWorks[i].author.length; j++) {

                        const renderAuthorHTML = (
                            <React.Fragment key={data.creativeWorks[i].author[j].id}>
                            <p>
                                {Artefacts.authorId}: {data.creativeWorks[i].author[j].id}<br/>
                                {File.authorName}: {data.creativeWorks[i].author[j].name}<br/>
                                {File.authorEMail}: {data.creativeWorks[i].author[j].email}<br/>
                                {File.authorURL}: {data.creativeWorks[i].author[j].url}<br/>
                            </p>
                            </React.Fragment>
                        )
                        artefactsInfo.push(renderAuthorHTML)
                    }
                 }
             }

              if( data.creativeWorks[i].copyrightHolder !== null ) {
                  if( typeof data.creativeWorks[i].copyrightHolder !== 'undefined' ) {
                      //console.log(data.creativeWorks[i].works)

                      for ( var j = 0; j < data.creativeWorks[i].copyrightHolder.length; j++) {

                          const renderCopyrightHTML = (
                              <React.Fragment key={data.creativeWorks[i].copyrightHolder[j].id}>
                              <p>
                                  {Artefacts.copyrightHolderId}: {data.creativeWorks[i].copyrightHolder[j].id}<br/>
                                  {File.copyrightHolderName}: {data.creativeWorks[i].copyrightHolder[j].name}<br/>
                                  {File.copyrightHolderEMail}: {data.creativeWorks[i].copyrightHolder[j].email}<br/>
                                  {File.copyrightHolderURL}: {data.creativeWorks[i].copyrightHolder[j].url}<br/>
                              </p>
                              </React.Fragment>
                          )
                          artefactsInfo.push(renderCopyrightHTML)
                      }
                   }
               }

               if( data.creativeWorks[i].publisher !== null ) {
                 if( typeof data.creativeWorks[i].publisher !== 'undefined' ) {

                    for ( var j = 0; j < data.creativeWorks[i].publisher.length; j++) {

                        const renderPublisherHTML = (
                           <React.Fragment key={data.creativeWorks[i].publisher[j].id}>
                           <p>
                               {Artefacts.publisherId}: {data.creativeWorks[i].publisher[j].id}<br/>
                               {File.publisherName}: {data.creativeWorks[i].publisher[j].name}<br/>
                               {File.publisherEMail}: {data.creativeWorks[i].publisher[j].email}<br/>
                               {File.publisherURL}: {data.creativeWorks[i].publisher[j].url}<br/>
                           </p>
                           </React.Fragment>
                        )
                        artefactsInfo.push(renderPublisherHTML)
                    }
                  }
                }
        }
    }

    return artefactsInfo
}

export const getArtefactHTML = (data: CreativeWorksProps, artefactId: string): any[] => {

    let artefactInfo = []

    if( typeof data.creativeWorks !== 'undefined' ) {
        for ( var i = 0; i < data.creativeWorks.length; i++) {

            const thisArtefactId = data.creativeWorks[i].id

            if (artefactId == thisArtefactId) {

                const renderArtefactsHTML = (
                    <React.Fragment key={data.creativeWorks[i].id}>
                    <p>
                        {Artefacts.id}: {data.creativeWorks[i].id}<br/>
                        {Artefacts.artefactName}: {data.creativeWorks[i].name}<br/>
                        {Artefacts.url}: {data.creativeWorks[i].url}<br/>
                        {Artefacts.description}: {data.creativeWorks[i].description}<br/>
                        {Artefacts.dateCreated}: {data.creativeWorks[i].dateCreated}<br/>
                        {Artefacts.dateModified}: {data.creativeWorks[i].dateModified}<br/>
                        {Artefacts.workType}: {FormHelpers.works[data.creativeWorks[i].workType]}<br/>
                        {Artefacts.license}: {FormHelpers.licenses[data.creativeWorks[i].license]}
                    </p>
                    </React.Fragment>
                )
                artefactInfo.push(renderArtefactsHTML)

                if( data.creativeWorks[i].author !== null ) {
                    if( typeof data.creativeWorks[i].author !== 'undefined' ) {
                        //console.log(data.creativeWorks[i].works)

                        for ( var j = 0; j < data.creativeWorks[i].author.length; j++) {

                            const renderAuthorHTML = (
                                <React.Fragment key={data.creativeWorks[i].author[j].id}>
                                <p>
                                    {Artefacts.authorId}: {data.creativeWorks[i].author[j].id}<br/>
                                    {File.authorName}: {data.creativeWorks[i].author[j].name}<br/>
                                    {File.authorEMail}: {data.creativeWorks[i].author[j].email}<br/>
                                    {File.authorURL}: {data.creativeWorks[i].author[j].url}<br/>
                                </p>
                                </React.Fragment>
                            )
                            artefactInfo.push(renderAuthorHTML)
                        }
                     }
                 }

                  if( data.creativeWorks[i].copyrightHolder !== null ) {
                      if( typeof data.creativeWorks[i].copyrightHolder !== 'undefined' ) {
                          //console.log(data.creativeWorks[i].works)

                          for ( var j = 0; j < data.creativeWorks[i].copyrightHolder.length; j++) {

                              const renderCopyrightHTML = (
                                  <React.Fragment key={data.creativeWorks[i].copyrightHolder[j].id}>
                                  <p>
                                      {Artefacts.copyrightHolderId}: {data.creativeWorks[i].copyrightHolder[j].id}<br/>
                                      {File.copyrightHolderName}: {data.creativeWorks[i].copyrightHolder[j].name}<br/>
                                      {File.copyrightHolderEMail}: {data.creativeWorks[i].copyrightHolder[j].email}<br/>
                                      {File.copyrightHolderURL}: {data.creativeWorks[i].copyrightHolder[j].url}<br/>
                                  </p>
                                  </React.Fragment>
                              )
                              artefactInfo.push(renderCopyrightHTML)
                          }
                       }
                   }

                   if( data.creativeWorks[i].publisher !== null ) {
                     if( typeof data.creativeWorks[i].publisher !== 'undefined' ) {

                        for ( var j = 0; j < data.creativeWorks[i].publisher.length; j++) {

                            const renderPublisherHTML = (
                               <React.Fragment key={data.creativeWorks[i].publisher[j].id}>
                               <p>
                                   {Artefacts.publisherId}: {data.creativeWorks[i].publisher[j].id}<br/>
                                   {File.publisherName}: {data.creativeWorks[i].publisher[j].name}<br/>
                                   {File.publisherEMail}: {data.creativeWorks[i].publisher[j].email}<br/>
                                   {File.publisherURL}: {data.creativeWorks[i].publisher[j].url}<br/>
                               </p>
                               </React.Fragment>
                            )
                            artefactInfo.push(renderPublisherHTML)
                        }
                      }
                    }
            }
        }
    }

    return artefactInfo
}
