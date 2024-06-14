package musictools

import "testing"

func TestParseRecentlyPlayed(t *testing.T) {

	testString := `{
  "recenttracks": {
    "track": [
      {
        "artist": {
          "mbid": "fd87acc7-e0a0-4a45-bc2a-d2ab5c10be68",
          "#text": "Fontaines D.C."
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "031fd41c-ef20-43ac-9f69-49e008226558",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Big",
        "url": "https://www.last.fm/music/Fontaines+D.C./_/Big",
        "date": {
          "uts": "1718387674",
          "#text": "14 Jun 2024, 17:54"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Justice"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "4b65b135-b72c-4d8e-a5a3-caf05e82f198",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Phantom Pt. II (Soulwax remix)",
        "url": "https://www.last.fm/music/Justice/_/Phantom+Pt.+II+(Soulwax+remix)",
        "date": {
          "uts": "1718387446",
          "#text": "14 Jun 2024, 17:50"
        }
      },
      {
        "artist": {
          "mbid": "9fab430a-f1af-4a69-bf68-6857a7a50a68",
          "#text": "Confidence Man"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "3d891037-ea03-47ba-aaa8-38e1dad3ed90",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Boyfriend",
        "url": "https://www.last.fm/music/Confidence+Man/_/Boyfriend",
        "date": {
          "uts": "1718387246",
          "#text": "14 Jun 2024, 17:47"
        }
      },
      {
        "artist": {
          "mbid": "74963434-dcb6-4b14-98cc-99873b06db66",
          "#text": "Future Islands"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "e381e936-0b69-43ae-ba79-d372c7597394",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Seasons (Waiting on You)",
        "url": "https://www.last.fm/music/Future+Islands/_/Seasons+(Waiting+on+You)",
        "date": {
          "uts": "1718386965",
          "#text": "14 Jun 2024, 17:42"
        }
      },
      {
        "artist": {
          "mbid": "7fb50287-029d-47cc-825a-235ca28024b2",
          "#text": "Soft Cell"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "73d3c54d-f5a9-39eb-b7af-fcceb8af131d",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Tainted Love",
        "url": "https://www.last.fm/music/Soft+Cell/_/Tainted+Love",
        "date": {
          "uts": "1718386816",
          "#text": "14 Jun 2024, 17:40"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Nilüfer Yanya"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Like I Say I Runaway",
        "url": "https://www.last.fm/music/Nil%C3%BCfer+Yanya/_/Like+I+Say+I+Runaway",
        "date": {
          "uts": "1718386639",
          "#text": "14 Jun 2024, 17:37"
        }
      },
      {
        "artist": {
          "mbid": "ea4dfa26-f633-4da6-a52a-f49ea4897b58",
          "#text": "R.E.M."
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "d4e0a590-e461-499a-8f62-b595b6c65ada",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Losing My Religion",
        "url": "https://www.last.fm/music/R.E.M./_/Losing+My+Religion",
        "date": {
          "uts": "1718386371",
          "#text": "14 Jun 2024, 17:32"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Travis"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Why Does It Always Rain on Me?",
        "url": "https://www.last.fm/music/Travis/_/Why+Does+It+Always+Rain+on+Me%3F",
        "date": {
          "uts": "1718385973",
          "#text": "14 Jun 2024, 17:26"
        }
      },
      {
        "artist": {
          "mbid": "21666405-49f8-44f6-94bb-e064f202f7ba",
          "#text": "The Waterboys"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "75d6cadc-ccaf-3c9c-be53-326f82385595",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Fisherman's Blues",
        "url": "https://www.last.fm/music/The+Waterboys/_/Fisherman%27s+Blues",
        "date": {
          "uts": "1718385730",
          "#text": "14 Jun 2024, 17:22"
        }
      },
      {
        "artist": {
          "mbid": "34c63966-445c-4613-afe1-4f0e1e53ae9a",
          "#text": "Fatboy Slim"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Praise You",
        "url": "https://www.last.fm/music/Fatboy+Slim/_/Praise+You",
        "date": {
          "uts": "1718385454",
          "#text": "14 Jun 2024, 17:17"
        }
      },
      {
        "artist": {
          "mbid": "4a8d9623-4d6c-4b7c-8dc5-5d5319ab8a20",
          "#text": "Jean Knight"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "43406c9e-87f1-4369-bf47-4ed7ea39a82f",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Mr Big Stuff",
        "url": "https://www.last.fm/music/Jean+Knight/_/Mr+Big+Stuff",
        "date": {
          "uts": "1718385309",
          "#text": "14 Jun 2024, 17:15"
        }
      },
      {
        "artist": {
          "mbid": "ec44206a-c3cc-419d-9ef9-08971c1297dd",
          "#text": "Coach Party"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "All I Wanna Do Is Hate",
        "url": "https://www.last.fm/music/Coach+Party/_/All+I+Wanna+Do+Is+Hate",
        "date": {
          "uts": "1718385112",
          "#text": "14 Jun 2024, 17:11"
        }
      },
      {
        "artist": {
          "mbid": "67f66c07-6e61-4026-ade5-7e782fad3a5d",
          "#text": "Foo Fighters"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Monkey Wrench",
        "url": "https://www.last.fm/music/Foo+Fighters/_/Monkey+Wrench",
        "date": {
          "uts": "1718384817",
          "#text": "14 Jun 2024, 17:06"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Shygirl"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "encore (feat. Danny L Harle)",
        "url": "https://www.last.fm/music/Shygirl/_/encore+(feat.+Danny+L+Harle)",
        "date": {
          "uts": "1718384662",
          "#text": "14 Jun 2024, 17:04"
        }
      },
      {
        "artist": {
          "mbid": "6e5c1366-5891-4e01-bb13-6c8b2c8cdd35",
          "#text": "MJ Cole"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "3ae9bcd2-ac67-46d1-a620-a243af4de2b6",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Crazy Love",
        "url": "https://www.last.fm/music/MJ+Cole/_/Crazy+Love",
        "date": {
          "uts": "1718384463",
          "#text": "14 Jun 2024, 17:01"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Mannequin P"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Loud Bark",
        "url": "https://www.last.fm/music/Mannequin+P/_/Loud+Bark",
        "date": {
          "uts": "1718384276",
          "#text": "14 Jun 2024, 16:57"
        }
      },
      {
        "artist": {
          "mbid": "329f95db-f02e-40bf-a5f7-0dea11529ce3",
          "#text": "Enter Shikari"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "07f8747f-97a9-3c2f-994b-90425af50f27",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Sssnakepit",
        "url": "https://www.last.fm/music/Enter+Shikari/_/Sssnakepit",
        "date": {
          "uts": "1718384015",
          "#text": "14 Jun 2024, 16:53"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Mungo’s Hi Fi"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "More Fyah (feat. Eva Lazarus)",
        "url": "https://www.last.fm/music/Mungo%E2%80%99s+Hi+Fi/_/More+Fyah+(feat.+Eva+Lazarus)",
        "date": {
          "uts": "1718383802",
          "#text": "14 Jun 2024, 16:50"
        }
      },
      {
        "artist": {
          "mbid": "85e67b0f-afd2-46c2-88b2-c3ba9b0883f2",
          "#text": "Swim Deep"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "0f569b86-60a2-4950-bc03-be065a56fa4f",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Glitter",
        "url": "https://www.last.fm/music/Swim+Deep/_/Glitter",
        "date": {
          "uts": "1718383542",
          "#text": "14 Jun 2024, 16:45"
        }
      },
      {
        "artist": {
          "mbid": "ccce2053-7007-4c36-b1e7-f8fcf5023a12",
          "#text": "Dexys Midnight Runners"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Geno",
        "url": "https://www.last.fm/music/Dexys+Midnight+Runners/_/Geno",
        "date": {
          "uts": "1718383339",
          "#text": "14 Jun 2024, 16:42"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Amyl and the Sniffers"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "U Should Not Be Doin That",
        "url": "https://www.last.fm/music/Amyl+and+the+Sniffers/_/U+Should+Not+Be+Doin+That",
        "date": {
          "uts": "1718383095",
          "#text": "14 Jun 2024, 16:38"
        }
      },
      {
        "artist": {
          "mbid": "2aaf7396-6ab8-40f3-9776-a41c42c8e26b",
          "#text": "LCD Soundsystem"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "0313bc3f-717d-40d7-a742-23f7bb89f375",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "tonite",
        "url": "https://www.last.fm/music/LCD+Soundsystem/_/tonite",
        "date": {
          "uts": "1718382775",
          "#text": "14 Jun 2024, 16:32"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Royel Otis"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Claw Foot",
        "url": "https://www.last.fm/music/Royel+Otis/_/Claw+Foot",
        "date": {
          "uts": "1718382389",
          "#text": "14 Jun 2024, 16:26"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Gretel"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Far Out",
        "url": "https://www.last.fm/music/Gretel/_/Far+Out",
        "date": {
          "uts": "1718382225",
          "#text": "14 Jun 2024, 16:23"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Angélica Garcia"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Gemini",
        "url": "https://www.last.fm/music/Ang%C3%A9lica+Garcia/_/Gemini",
        "date": {
          "uts": "1718382019",
          "#text": "14 Jun 2024, 16:20"
        }
      },
      {
        "artist": {
          "mbid": "54b0584e-4e5e-4976-80d4-eea7a82a7213",
          "#text": "Electric Six"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Danger! High Voltage!",
        "url": "https://www.last.fm/music/Electric+Six/_/Danger!+High+Voltage!",
        "date": {
          "uts": "1718381768",
          "#text": "14 Jun 2024, 16:16"
        }
      },
      {
        "artist": {
          "mbid": "9b085502-2c93-446f-ab1b-858e281471ff",
          "#text": "Riton & Kah‐Lo"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "4b6599b0-ffde-43d9-88bc-1997b8bd07d7",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Fake ID",
        "url": "https://www.last.fm/music/Riton+&+Kah%E2%80%90Lo/_/Fake+ID",
        "date": {
          "uts": "1718381538",
          "#text": "14 Jun 2024, 16:12"
        }
      },
      {
        "artist": {
          "mbid": "d350bc0e-10bb-47f1-9027-d1f011ae9aa8",
          "#text": "Au Pairs"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "07a099c8-d5f4-32ac-935a-e6432e9e60cf",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Inconvenience",
        "url": "https://www.last.fm/music/Au+Pairs/_/Inconvenience",
        "date": {
          "uts": "1718381339",
          "#text": "14 Jun 2024, 16:08"
        }
      },
      {
        "artist": {
          "mbid": "892500b7-09a6-4049-ba92-2d192dd70563",
          "#text": "Biffy Clyro"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "0647e687-6615-44b0-ac14-a31f8741eb9b",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Mountains",
        "url": "https://www.last.fm/music/Biffy+Clyro/_/Mountains",
        "date": {
          "uts": "1718381137",
          "#text": "14 Jun 2024, 16:05"
        }
      },
      {
        "artist": {
          "mbid": "4b585938-f271-45e2-b19a-91c634b5e396",
          "#text": "Kate Bush"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Running Up That Hill (A Deal with God)",
        "url": "https://www.last.fm/music/Kate+Bush/_/Running+Up+That+Hill+(A+Deal+with+God)",
        "date": {
          "uts": "1718380851",
          "#text": "14 Jun 2024, 16:00"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "86TVs"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Someone Else's Dream",
        "url": "https://www.last.fm/music/86TVs/_/Someone+Else%27s+Dream",
        "date": {
          "uts": "1718380673",
          "#text": "14 Jun 2024, 15:57"
        }
      },
      {
        "artist": {
          "mbid": "e520459c-dff4-491d-a6e4-c97be35e0044",
          "#text": "Frank Ocean"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "dddbb3b9-7ed5-362c-983f-0d28b047dbf0",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Novacane",
        "url": "https://www.last.fm/music/Frank+Ocean/_/Novacane",
        "date": {
          "uts": "1718380300",
          "#text": "14 Jun 2024, 15:51"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Bodega"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "14792d5e-6741-4818-be40-74f41b4ff11e",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "ATM",
        "url": "https://www.last.fm/music/Bodega/_/ATM",
        "date": {
          "uts": "1718380142",
          "#text": "14 Jun 2024, 15:49"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Robin Hall & Jimmy MacGregor"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Football Crazy",
        "url": "https://www.last.fm/music/Robin+Hall+&+Jimmy+MacGregor/_/Football+Crazy",
        "date": {
          "uts": "1718379967",
          "#text": "14 Jun 2024, 15:46"
        }
      },
      {
        "artist": {
          "mbid": "2ac533ad-f3ec-4083-a2d4-6f30393414b3",
          "#text": "CMAT"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "31e895af-cb29-41e8-8186-5e29631a50b6",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Aw, Shoot!",
        "url": "https://www.last.fm/music/CMAT/_/Aw,+Shoot!",
        "date": {
          "uts": "1718379429",
          "#text": "14 Jun 2024, 15:37"
        }
      },
      {
        "artist": {
          "mbid": "4d2956d1-a3f7-44bb-9a41-67563e1a0c94",
          "#text": "Blondie"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "One Way or Another",
        "url": "https://www.last.fm/music/Blondie/_/One+Way+or+Another",
        "date": {
          "uts": "1718379229",
          "#text": "14 Jun 2024, 15:33"
        }
      },
      {
        "artist": {
          "mbid": "4eeea7eb-7554-48cb-9e50-c26d434ed9ef",
          "#text": "Lambrini Girls"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "af621b69-febe-4c31-b2e1-117fae55a1be",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Body of Mine",
        "url": "https://www.last.fm/music/Lambrini+Girls/_/Body+of+Mine",
        "date": {
          "uts": "1718378892",
          "#text": "14 Jun 2024, 15:28"
        }
      },
      {
        "artist": {
          "mbid": "d5da1841-9bc8-4813-9f89-11098090148e",
          "#text": "The Fall"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "e1349bb1-13ee-4064-9339-02881c6488d1",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Totally Wired",
        "url": "https://www.last.fm/music/The+Fall/_/Totally+Wired",
        "date": {
          "uts": "1718378687",
          "#text": "14 Jun 2024, 15:24"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Fulu Miziki"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Mosala",
        "url": "https://www.last.fm/music/Fulu+Miziki/_/Mosala",
        "date": {
          "uts": "1718378436",
          "#text": "14 Jun 2024, 15:20"
        }
      },
      {
        "artist": {
          "mbid": "095b2041-4975-4ba3-a92e-53fa3459107f",
          "#text": "Catatonia"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Road Rage",
        "url": "https://www.last.fm/music/Catatonia/_/Road+Rage",
        "date": {
          "uts": "1718378186",
          "#text": "14 Jun 2024, 15:16"
        }
      },
      {
        "artist": {
          "mbid": "34e56f20-ddda-48c9-a4bb-6bb12df57de0",
          "#text": "The Style Council"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Shout To The Top",
        "url": "https://www.last.fm/music/The+Style+Council/_/Shout+To+The+Top",
        "date": {
          "uts": "1718377951",
          "#text": "14 Jun 2024, 15:12"
        }
      },
      {
        "artist": {
          "mbid": "8ee9e6fe-a629-4ff9-a4d5-ecb79b637350",
          "#text": "Lava La Rue"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "LOVEBITES",
        "url": "https://www.last.fm/music/Lava+La+Rue/_/LOVEBITES",
        "date": {
          "uts": "1718377749",
          "#text": "14 Jun 2024, 15:09"
        }
      },
      {
        "artist": {
          "mbid": "3cdb40fe-a63e-4bb9-b40d-17cda5f50979",
          "#text": "Little Simz"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "06cae4bc-52e4-4b57-9c9c-682da38a3c54",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Introvert",
        "url": "https://www.last.fm/music/Little+Simz/_/Introvert",
        "date": {
          "uts": "1718377391",
          "#text": "14 Jun 2024, 15:03"
        }
      },
      {
        "artist": {
          "mbid": "3cdb40fe-a63e-4bb9-b40d-17cda5f50979",
          "#text": "Little Simz"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "06cae4bc-52e4-4b57-9c9c-682da38a3c54",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Introvert",
        "url": "https://www.last.fm/music/Little+Simz/_/Introvert",
        "date": {
          "uts": "1718377312",
          "#text": "14 Jun 2024, 15:01"
        }
      },
      {
        "artist": {
          "mbid": "63aa26c3-d59b-4da4-84ac-716b54f1ef4d",
          "#text": "Tame Impala"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "7d727f9d-c3ed-4b9b-9646-c4d1d6025624",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Elephant",
        "url": "https://www.last.fm/music/Tame+Impala/_/Elephant",
        "date": {
          "uts": "1718377214",
          "#text": "14 Jun 2024, 15:00"
        }
      },
      {
        "artist": {
          "mbid": "1303b976-b862-4f04-94fd-a8d444e06714",
          "#text": "The Proclaimers"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "12b0edb9-0441-331f-a095-811f51a36a35",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Sunshine On Leith",
        "url": "https://www.last.fm/music/The+Proclaimers/_/Sunshine+On+Leith",
        "date": {
          "uts": "1718376893",
          "#text": "14 Jun 2024, 14:54"
        }
      },
      {
        "artist": {
          "mbid": "5588744b-0032-48d2-9ea5-ac74defefcdd",
          "#text": "Galliano"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Circles Going Round The Sun",
        "url": "https://www.last.fm/music/Galliano/_/Circles+Going+Round+The+Sun",
        "date": {
          "uts": "1718376631",
          "#text": "14 Jun 2024, 14:50"
        }
      },
      {
        "artist": {
          "mbid": "27ca9c68-650a-4659-b8ef-013e5698ff55",
          "#text": "John Grant"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "All That School For Nothing",
        "url": "https://www.last.fm/music/John+Grant/_/All+That+School+For+Nothing",
        "date": {
          "uts": "1718376358",
          "#text": "14 Jun 2024, 14:45"
        }
      },
      {
        "artist": {
          "mbid": "42dd1862-7255-47ed-8584-651abce3bf53",
          "#text": "Nubya Garcia"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "84a3b4c2-d179-4090-87ce-1213cae7b874",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "The Seer",
        "url": "https://www.last.fm/music/Nubya+Garcia/_/The+Seer",
        "date": {
          "uts": "1718376124",
          "#text": "14 Jun 2024, 14:42"
        }
      },
      {
        "artist": {
          "mbid": "",
          "#text": "Fat Dog"
        },
        "streamable": "0",
        "image": [
          {
            "size": "small",
            "#text": ""
          },
          {
            "size": "medium",
            "#text": ""
          },
          {
            "size": "large",
            "#text": ""
          },
          {
            "size": "extralarge",
            "#text": ""
          }
        ],
        "mbid": "",
        "album": {
          "mbid": "",
          "#text": ""
        },
        "name": "Running",
        "url": "https://www.last.fm/music/Fat+Dog/_/Running",
        "date": {
          "uts": "1718375897",
          "#text": "14 Jun 2024, 14:38"
        }
      }
    ],
    "@attr": {
      "user": "bbc6music",
      "totalPages": "31480",
      "page": "1",
      "perPage": "50",
      "total": "1573985"
    }
  }
}`

	responseBytes := []byte(testString)

	parsedResponse := parseRecentlyPlayed(responseBytes)

	expected := "Big"
	got := parsedResponse.Recenttracks.Track[0].Name

	if expected != got {
		t.Errorf("expected song name '%v, got %v", expected, got)
	}

}
