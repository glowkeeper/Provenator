type NumberOptionType = { label: string; value: number }
type StringOptionType = { label: string; value: string }

class FormHelpers {

    static readonly works = [
        "NONE",
        "Article",
        "Book",
        "Blog",
        "Clip",
        "Digital Document",
        "Drawing",
        "Movie",
        "Music Composition",
        "Music Recording",
        "Painting",
        "Photograph",
        "Poster",
        "Short Story",
        "Software Application",
        "Software Source Code",
        "Thesis",
        "Visual Artwork",
        "Web Content",
        "Webpage",
        "Website",
        "MAX"
    ]

    static readonly worksSelect = [
        {value: 1, label: `${FormHelpers.works[1]}`},
        {value: 2, label: `${FormHelpers.works[2]}`},
        {value: 3, label: `${FormHelpers.works[3]}`},
        {value: 4, label: `${FormHelpers.works[4]}`},
        {value: 5, label: `${FormHelpers.works[5]}`},
        {value: 6, label: `${FormHelpers.works[6]}`},
        {value: 7, label: `${FormHelpers.works[7]}`},
        {value: 8, label: `${FormHelpers.works[8]}`},
        {value: 9, label: `${FormHelpers.works[9]}`},
        {value: 10, label: `${FormHelpers.works[10]}`},
        {value: 11, label: `${FormHelpers.works[11]}`},
        {value: 12, label: `${FormHelpers.works[12]}`},
        {value: 13, label: `${FormHelpers.works[13]}`},
        {value: 14, label: `${FormHelpers.works[14]}`},
        {value: 15, label: `${FormHelpers.works[15]}`},
        {value: 16, label: `${FormHelpers.works[16]}`},
        {value: 17, label: `${FormHelpers.works[17]}`},
        {value: 18, label: `${FormHelpers.works[18]}`},
        {value: 19, label: `${FormHelpers.works[19]}`},
        {value: 20, label: `${FormHelpers.works[20]}`}
    ]

    static readonly licenses = [
        "None",
        "CC BY (Attribution)",
        "CC BY-SA (Attribution-Share-Alike)",
        "CC BY-ND (Attribution-No-Derivatives)",
        "CC BY-NC (Attribution-Non-Commercial)",
        "CC BY-NC-SA (Attribution-NonCommercial-ShareAlike)",
        "CC BY-NC-ND (Attribution-NonCommercial-ShareAlike-NoDerivatives)",
        "Apache License 2.0",
        "GNU General Public License v3.0",
        "MIT License",
        "BSD 2-Clause Simplified License",
        "BSD 3-Clause New or Revised License",
        "Boost Software License 1.0",
        "Creative Commons Zero v1.0 Universal",
        "Eclipse Public License 2.0",
        "GNU Affero General Public License v3.0",
        "GNU General Public License v2.0",
        "GNU Lesser General Public License v2.1",
        "Mozilla Public License 2.0",
        "The Unlicense"
    ]

    static readonly licensesSelect = [
        {value: 0, label: `${FormHelpers.licenses[0]}`},
        {value: 1, label: `${FormHelpers.licenses[1]}`},
        {value: 2, label: `${FormHelpers.licenses[2]}`},
        {value: 3, label: `${FormHelpers.licenses[3]}`},
        {value: 4, label: `${FormHelpers.licenses[4]}`},
        {value: 5, label: `${FormHelpers.licenses[5]}`},
        {value: 6, label: `${FormHelpers.licenses[6]}`},
        {value: 7, label: `${FormHelpers.licenses[7]}`},
        {value: 8, label: `${FormHelpers.licenses[8]}`},
        {value: 9, label: `${FormHelpers.licenses[9]}`},
        {value: 10, label: `${FormHelpers.licenses[10]}`},
        {value: 11, label: `${FormHelpers.licenses[11]}`},
        {value: 12, label: `${FormHelpers.licenses[12]}`},
        {value: 13, label: `${FormHelpers.licenses[13]}`},
        {value: 14, label: `${FormHelpers.licenses[14]}`},
        {value: 15, label: `${FormHelpers.licenses[15]}`},
        {value: 16, label: `${FormHelpers.licenses[16]}`},
        {value: 17, label: `${FormHelpers.licenses[17]}`},
        {value: 18, label: `${FormHelpers.licenses[18]}`},        
        {value: 19, label: `${FormHelpers.licenses[19]}`}
    ]
}

export { NumberOptionType, StringOptionType, FormHelpers }
