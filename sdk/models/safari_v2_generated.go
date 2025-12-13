package models

import (
	"time"
)

type SafariV2FeedResponse struct {
	AppVersions []struct {
		AppVersion string `json:"AppVersion"`
		Latest     struct {
			ActivelyExploitedCvEs []string `json:"ActivelyExploitedCVEs"`
			CvEs                  struct {
				Cve202440776 *struct{} `json:"CVE-2024-40776,omitempty"`
				Cve202440779 *struct{} `json:"CVE-2024-40779,omitempty"`
				Cve202440780 *struct{} `json:"CVE-2024-40780,omitempty"`
				Cve202440782 *struct{} `json:"CVE-2024-40782,omitempty"`
				Cve202440785 *struct{} `json:"CVE-2024-40785,omitempty"`
				Cve202440789 *struct{} `json:"CVE-2024-40789,omitempty"`
				Cve202440794 *struct{} `json:"CVE-2024-40794,omitempty"`
				Cve202440817 *struct{} `json:"CVE-2024-40817,omitempty"`
				Cve202444185 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44185,omitempty"`
				Cve202444206 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44206,omitempty"`
				Cve20244558  *struct{} `json:"CVE-2024-4558,omitempty"`
				Cve202454551 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-54551,omitempty"`
				Cve20248906  *struct{} `json:"CVE-2024-8906,omitempty"`
				Cve202514174 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Remote         bool     `json:"remote"`
					Severity       string   `json:"severity"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-14174,omitempty"`
				Cve202524188 *struct{} `json:"CVE-2025-24188,omitempty"`
				Cve202531273 *struct{} `json:"CVE-2025-31273,omitempty"`
				Cve202531277 *struct{} `json:"CVE-2025-31277,omitempty"`
				Cve202531278 *struct{} `json:"CVE-2025-31278,omitempty"`
				Cve202543211 *struct{} `json:"CVE-2025-43211,omitempty"`
				Cve202543212 *struct{} `json:"CVE-2025-43212,omitempty"`
				Cve202543213 *struct{} `json:"CVE-2025-43213,omitempty"`
				Cve202543214 *struct{} `json:"CVE-2025-43214,omitempty"`
				Cve202543216 *struct{} `json:"CVE-2025-43216,omitempty"`
				Cve202543227 *struct{} `json:"CVE-2025-43227,omitempty"`
				Cve202543228 *struct{} `json:"CVE-2025-43228,omitempty"`
				Cve202543229 *struct{} `json:"CVE-2025-43229,omitempty"`
				Cve202543240 *struct{} `json:"CVE-2025-43240,omitempty"`
				Cve202543265 *struct{} `json:"CVE-2025-43265,omitempty"`
				Cve202543501 *struct{} `json:"CVE-2025-43501,omitempty"`
				Cve202543511 *struct{} `json:"CVE-2025-43511,omitempty"`
				Cve202543526 *struct{} `json:"CVE-2025-43526,omitempty"`
				Cve202543529 *struct{} `json:"CVE-2025-43529,omitempty"`
				Cve202543531 *struct{} `json:"CVE-2025-43531,omitempty"`
				Cve202543535 *struct{} `json:"CVE-2025-43535,omitempty"`
				Cve202543536 *struct{} `json:"CVE-2025-43536,omitempty"`
				Cve202543541 *struct{} `json:"CVE-2025-43541,omitempty"`
				Cve202546282 *struct{} `json:"CVE-2025-46282,omitempty"`
				Cve20256558  *struct {
					Exploited      bool   `json:"exploited"`
					Kev            bool   `json:"kev"`
					MatchedPattern string `json:"matched_pattern"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2025-6558,omitempty"`
				Cve20257424 *struct{} `json:"CVE-2025-7424,omitempty"`
				Cve20257425 *struct{} `json:"CVE-2025-7425,omitempty"`
			} `json:"CVEs"`
			DaysSincePreviousRelease int       `json:"DaysSincePreviousRelease"`
			ProductName              string    `json:"ProductName"`
			ProductVersion           string    `json:"ProductVersion"`
			ReleaseDate              time.Time `json:"ReleaseDate"`
			ReleaseType              string    `json:"ReleaseType"`
			SecurityInfo             string    `json:"SecurityInfo"`
			UniqueCvEsCount          int       `json:"UniqueCVEsCount"`
			UpdateName               string    `json:"UpdateName"`
		} `json:"Latest"`
		SecurityReleases []struct {
			ActivelyExploitedCvEs []string `json:"ActivelyExploitedCVEs"`
			CvEs                  struct {
				Cve202335074 *struct{} `json:"CVE-2023-35074,omitempty"`
				Cve202339434 *struct{} `json:"CVE-2023-39434,omitempty"`
				Cve202340385 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-40385,omitempty"`
				Cve202340414 *struct{} `json:"CVE-2023-40414,omitempty"`
				Cve202340417 *struct{} `json:"CVE-2023-40417,omitempty"`
				Cve202340447 *struct{} `json:"CVE-2023-40447,omitempty"`
				Cve202340451 *struct{} `json:"CVE-2023-40451,omitempty"`
				Cve202341074 *struct{} `json:"CVE-2023-41074,omitempty"`
				Cve202341976 *struct{} `json:"CVE-2023-41976,omitempty"`
				Cve202341983 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-41983,omitempty"`
				Cve202341993 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-41993,omitempty"`
				Cve202342833 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42833,omitempty"`
				Cve202342843 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42843,omitempty"`
				Cve202342852 *struct{} `json:"CVE-2023-42852,omitempty"`
				Cve202342875 *struct{} `json:"CVE-2023-42875,omitempty"`
				Cve202342883 *struct{} `json:"CVE-2023-42883,omitempty"`
				Cve202342890 *struct{} `json:"CVE-2023-42890,omitempty"`
				Cve202342916 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42916,omitempty"`
				Cve202342917 *struct {
					Exploited      bool   `json:"exploited"`
					Kev            bool   `json:"kev"`
					MatchedPattern string `json:"matched_pattern"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2023-42917,omitempty"`
				Cve202342950 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42950,omitempty"`
				Cve202342956 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42956,omitempty"`
				Cve202342970 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2023-42970,omitempty"`
				Cve20241580 *struct {
					Exploited bool   `json:"exploited"`
					NistURL   string `json:"nist_url"`
				} `json:"CVE-2024-1580,omitempty"`
				Cve202423206 *struct{} `json:"CVE-2024-23206,omitempty"`
				Cve202423211 *struct{} `json:"CVE-2024-23211,omitempty"`
				Cve202423213 *struct{} `json:"CVE-2024-23213,omitempty"`
				Cve202423222 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-23222,omitempty"`
				Cve202423254 *struct{} `json:"CVE-2024-23254,omitempty"`
				Cve202423263 *struct{} `json:"CVE-2024-23263,omitempty"`
				Cve202423271 *struct {
					EntryAddedDate string `json:"entry_added_date"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2024-23271,omitempty"`
				Cve202423273 *struct{} `json:"CVE-2024-23273,omitempty"`
				Cve202423280 *struct{} `json:"CVE-2024-23280,omitempty"`
				Cve202423284 *struct {
					EntryAddedDate string `json:"entry_added_date"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2024-23284,omitempty"`
				Cve202427808 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27808,omitempty"`
				Cve202427820 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27820,omitempty"`
				Cve202427830 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27830,omitempty"`
				Cve202427833 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27833,omitempty"`
				Cve202427834 *struct{} `json:"CVE-2024-27834,omitempty"`
				Cve202427838 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27838,omitempty"`
				Cve202427844 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27844,omitempty"`
				Cve202427850 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27850,omitempty"`
				Cve202427851 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27851,omitempty"`
				Cve202427856 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-27856,omitempty"`
				Cve202440776 *struct{} `json:"CVE-2024-40776,omitempty"`
				Cve202440779 *struct{} `json:"CVE-2024-40779,omitempty"`
				Cve202440780 *struct{} `json:"CVE-2024-40780,omitempty"`
				Cve202440782 *struct{} `json:"CVE-2024-40782,omitempty"`
				Cve202440785 *struct{} `json:"CVE-2024-40785,omitempty"`
				Cve202440789 *struct{} `json:"CVE-2024-40789,omitempty"`
				Cve202440794 *struct{} `json:"CVE-2024-40794,omitempty"`
				Cve202440817 *struct{} `json:"CVE-2024-40817,omitempty"`
				Cve202440857 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-40857,omitempty"`
				Cve202440866 *struct{} `json:"CVE-2024-40866,omitempty"`
				Cve202444155 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44155,omitempty"`
				Cve202444185 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44185,omitempty"`
				Cve202444187 *struct{} `json:"CVE-2024-44187,omitempty"`
				Cve202444192 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44192,omitempty"`
				Cve202444202 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44202,omitempty"`
				Cve202444206 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44206,omitempty"`
				Cve202444212 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44212,omitempty"`
				Cve202444229 *struct{} `json:"CVE-2024-44229,omitempty"`
				Cve202444244 *struct{} `json:"CVE-2024-44244,omitempty"`
				Cve202444246 *struct{} `json:"CVE-2024-44246,omitempty"`
				Cve202444259 *struct{} `json:"CVE-2024-44259,omitempty"`
				Cve202444296 *struct{} `json:"CVE-2024-44296,omitempty"`
				Cve202444308 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-44308,omitempty"`
				Cve202444309 *struct {
					Exploited      bool   `json:"exploited"`
					Kev            bool   `json:"kev"`
					MatchedPattern string `json:"matched_pattern"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2024-44309,omitempty"`
				Cve20244558  *struct{} `json:"CVE-2024-4558,omitempty"`
				Cve202454467 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-54467,omitempty"`
				Cve202454479 *struct{} `json:"CVE-2024-54479,omitempty"`
				Cve202454502 *struct{} `json:"CVE-2024-54502,omitempty"`
				Cve202454505 *struct{} `json:"CVE-2024-54505,omitempty"`
				Cve202454508 *struct{} `json:"CVE-2024-54508,omitempty"`
				Cve202454534 *struct{} `json:"CVE-2024-54534,omitempty"`
				Cve202454542 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-54542,omitempty"`
				Cve202454543 *struct{} `json:"CVE-2024-54543,omitempty"`
				Cve202454551 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-54551,omitempty"`
				Cve202454658 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2024-54658,omitempty"`
				Cve20248906  *struct{} `json:"CVE-2024-8906,omitempty"`
				Cve202514174 *struct {
					Exploited      bool     `json:"exploited"`
					Kev            bool     `json:"kev"`
					MatchedPattern string   `json:"matched_pattern"`
					NistURL        string   `json:"nist_url"`
					Remote         bool     `json:"remote"`
					Severity       string   `json:"severity"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-14174,omitempty"`
				Cve202524113 *struct{} `json:"CVE-2025-24113,omitempty"`
				Cve202524128 *struct{} `json:"CVE-2025-24128,omitempty"`
				Cve202524143 *struct{} `json:"CVE-2025-24143,omitempty"`
				Cve202524150 *struct{} `json:"CVE-2025-24150,omitempty"`
				Cve202524158 *struct{} `json:"CVE-2025-24158,omitempty"`
				Cve202524162 *struct{} `json:"CVE-2025-24162,omitempty"`
				Cve202524167 *struct{} `json:"CVE-2025-24167,omitempty"`
				Cve202524169 *struct{} `json:"CVE-2025-24169,omitempty"`
				Cve202524180 *struct{} `json:"CVE-2025-24180,omitempty"`
				Cve202524188 *struct{} `json:"CVE-2025-24188,omitempty"`
				Cve202524189 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-24189,omitempty"`
				Cve202524192 *struct{} `json:"CVE-2025-24192,omitempty"`
				Cve202524201 *struct {
					Exploited      bool   `json:"exploited"`
					Kev            bool   `json:"kev"`
					MatchedPattern string `json:"matched_pattern"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2025-24201,omitempty"`
				Cve202524208 *struct{} `json:"CVE-2025-24208,omitempty"`
				Cve202524209 *struct{} `json:"CVE-2025-24209,omitempty"`
				Cve202524213 *struct{} `json:"CVE-2025-24213,omitempty"`
				Cve202524216 *struct{} `json:"CVE-2025-24216,omitempty"`
				Cve202524223 *struct{} `json:"CVE-2025-24223,omitempty"`
				Cve202524264 *struct{} `json:"CVE-2025-24264,omitempty"`
				Cve202530425 *struct{} `json:"CVE-2025-30425,omitempty"`
				Cve202530427 *struct{} `json:"CVE-2025-30427,omitempty"`
				Cve202530466 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-30466,omitempty"`
				Cve202530467 *struct{} `json:"CVE-2025-30467,omitempty"`
				Cve202531184 *struct{} `json:"CVE-2025-31184,omitempty"`
				Cve202531192 *struct{} `json:"CVE-2025-31192,omitempty"`
				Cve202531204 *struct{} `json:"CVE-2025-31204,omitempty"`
				Cve202531205 *struct{} `json:"CVE-2025-31205,omitempty"`
				Cve202531206 *struct{} `json:"CVE-2025-31206,omitempty"`
				Cve202531215 *struct{} `json:"CVE-2025-31215,omitempty"`
				Cve202531217 *struct{} `json:"CVE-2025-31217,omitempty"`
				Cve202531223 *struct{} `json:"CVE-2025-31223,omitempty"`
				Cve202531238 *struct{} `json:"CVE-2025-31238,omitempty"`
				Cve202531254 *struct{} `json:"CVE-2025-31254,omitempty"`
				Cve202531257 *struct{} `json:"CVE-2025-31257,omitempty"`
				Cve202531266 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-31266,omitempty"`
				Cve202531273 *struct{} `json:"CVE-2025-31273,omitempty"`
				Cve202531277 *struct{} `json:"CVE-2025-31277,omitempty"`
				Cve202531278 *struct{} `json:"CVE-2025-31278,omitempty"`
				Cve202543211 *struct{} `json:"CVE-2025-43211,omitempty"`
				Cve202543212 *struct{} `json:"CVE-2025-43212,omitempty"`
				Cve202543213 *struct{} `json:"CVE-2025-43213,omitempty"`
				Cve202543214 *struct{} `json:"CVE-2025-43214,omitempty"`
				Cve202543216 *struct{} `json:"CVE-2025-43216,omitempty"`
				Cve202543227 *struct{} `json:"CVE-2025-43227,omitempty"`
				Cve202543228 *struct{} `json:"CVE-2025-43228,omitempty"`
				Cve202543229 *struct{} `json:"CVE-2025-43229,omitempty"`
				Cve202543240 *struct{} `json:"CVE-2025-43240,omitempty"`
				Cve202543265 *struct{} `json:"CVE-2025-43265,omitempty"`
				Cve202543272 *struct{} `json:"CVE-2025-43272,omitempty"`
				Cve202543327 *struct{} `json:"CVE-2025-43327,omitempty"`
				Cve202543342 *struct{} `json:"CVE-2025-43342,omitempty"`
				Cve202543343 *struct{} `json:"CVE-2025-43343,omitempty"`
				Cve202543356 *struct{} `json:"CVE-2025-43356,omitempty"`
				Cve202543368 *struct{} `json:"CVE-2025-43368,omitempty"`
				Cve202543376 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-43376,omitempty"`
				Cve202543392 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43392,omitempty"`
				Cve202543419 *struct {
					EntryAddedDate string   `json:"entry_added_date"`
					NistURL        string   `json:"nist_url"`
					Tags           []string `json:"tags"`
				} `json:"CVE-2025-43419,omitempty"`
				Cve202543421 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43421,omitempty"`
				Cve202543425 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43425,omitempty"`
				Cve202543427 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43427,omitempty"`
				Cve202543429 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43429,omitempty"`
				Cve202543430 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43430,omitempty"`
				Cve202543431 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43431,omitempty"`
				Cve202543432 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43432,omitempty"`
				Cve202543433 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43433,omitempty"`
				Cve202543434 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43434,omitempty"`
				Cve202543435 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43435,omitempty"`
				Cve202543438 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43438,omitempty"`
				Cve202543440 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43440,omitempty"`
				Cve202543441 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43441,omitempty"`
				Cve202543443 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43443,omitempty"`
				Cve202543457 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43457,omitempty"`
				Cve202543458 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43458,omitempty"`
				Cve202543480 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43480,omitempty"`
				Cve202543493 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43493,omitempty"`
				Cve202543501 *struct{} `json:"CVE-2025-43501,omitempty"`
				Cve202543502 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43502,omitempty"`
				Cve202543503 *struct {
					Exploited bool     `json:"exploited"`
					NistURL   string   `json:"nist_url"`
					Tags      []string `json:"tags"`
				} `json:"CVE-2025-43503,omitempty"`
				Cve202543511 *struct{} `json:"CVE-2025-43511,omitempty"`
				Cve202543526 *struct{} `json:"CVE-2025-43526,omitempty"`
				Cve202543529 *struct{} `json:"CVE-2025-43529,omitempty"`
				Cve202543531 *struct{} `json:"CVE-2025-43531,omitempty"`
				Cve202543535 *struct{} `json:"CVE-2025-43535,omitempty"`
				Cve202543536 *struct{} `json:"CVE-2025-43536,omitempty"`
				Cve202543541 *struct{} `json:"CVE-2025-43541,omitempty"`
				Cve202546282 *struct{} `json:"CVE-2025-46282,omitempty"`
				Cve20256558  *struct {
					Exploited      bool   `json:"exploited"`
					Kev            bool   `json:"kev"`
					MatchedPattern string `json:"matched_pattern"`
					NistURL        string `json:"nist_url"`
				} `json:"CVE-2025-6558,omitempty"`
				Cve20257424 *struct{} `json:"CVE-2025-7424,omitempty"`
				Cve20257425 *struct{} `json:"CVE-2025-7425,omitempty"`
			} `json:"CVEs"`
			DaysSincePreviousRelease int       `json:"DaysSincePreviousRelease,omitempty"`
			ProductName              string    `json:"ProductName"`
			ProductVersion           string    `json:"ProductVersion"`
			ReleaseDate              time.Time `json:"ReleaseDate"`
			ReleaseType              string    `json:"ReleaseType"`
			SecurityInfo             string    `json:"SecurityInfo"`
			UniqueCvEsCount          int       `json:"UniqueCVEsCount"`
			UpdateName               string    `json:"UpdateName"`
		} `json:"SecurityReleases"`
	} `json:"AppVersions"`
	UpdateHash string `json:"UpdateHash"`
	Version    string `json:"Version"`
}
