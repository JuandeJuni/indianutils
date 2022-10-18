package utils

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/JuandeJuni/discordhooks"
)

func GetKw() []string {
	Kwlist, _ := ReadLines("kw.txt")
	return Kwlist
}

type VegnoProd []struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func HookStockSizes(title string, link string, image string, price string, description string, sizestring string, stockstring string, color int, footer string, hook string) {
	embed := discordhooks.Embed{
		Title:       title,
		Url:         link,
		Description: description,
		Color:       color,
		Thumbnail:   discordhooks.Thumbnail{Url: image},
		Fields:      []discordhooks.Field{{Name: "Price :money_with_wings: ", Value: price, Inline: true}, {Name: "**Sizes** :athletic_shoe: ", Value: sizestring, Inline: true}, {Name: "**Stock** :shopping_cart: ", Value: stockstring, Inline: true}},
		Footer:      discordhooks.Footer{Text: footer, Icon_url: "https://media.discordapp.net/attachments/994607849730748497/1014883850331095120/panda_discohook-696969696969-resized.jpg"},
		Timestamp:   time.Now(),
	}
	discordhooks.SendEmbed(hook, embed)
	discordhooks.SendEmbed("https://discord.com/api/webhooks/1011985224508395531/2-WCpEULIO8uBAmXuUte3THKMAljlWguqKKS1bM6IWMdY1iaI_BKZX852CamAVYh9zAp", embed)
}
func HookSizes(title string, link string, image string, price string, description string, sizestring string, color int, footer string, hook string) {
	embed := discordhooks.Embed{
		Title:       title,
		Url:         link,
		Description: description,
		Color:       color,
		Thumbnail:   discordhooks.Thumbnail{Url: image},
		Fields:      []discordhooks.Field{{Name: "Price :money_with_wings: ", Value: price, Inline: true}, {Name: "**Sizes** :athletic_shoe: ", Value: sizestring, Inline: true}},
		Footer:      discordhooks.Footer{Text: footer, Icon_url: "https://media.discordapp.net/attachments/994607849730748497/1014883850331095120/panda_discohook-696969696969-resized.jpg"},
		Timestamp:   time.Now(),
	}
	discordhooks.SendEmbed(hook, embed)
	discordhooks.SendEmbed("https://discord.com/api/webhooks/1011985224508395531/2-WCpEULIO8uBAmXuUte3THKMAljlWguqKKS1bM6IWMdY1iaI_BKZX852CamAVYh9zAp", embed)
}
func Hook(title string, link string, image string, price string, description string, color int, footer string, hook string) {
	embed := discordhooks.Embed{
		Title:       title,
		Url:         link,
		Description: description,
		Color:       color,
		Thumbnail:   discordhooks.Thumbnail{Url: image},
		Fields:      []discordhooks.Field{{Name: "Price :money_with_wings: ", Value: price, Inline: true}},
		Footer:      discordhooks.Footer{Text: footer, Icon_url: "https://media.discordapp.net/attachments/994607849730748497/1014883850331095120/panda_discohook-696969696969-resized.jpg"},
		Timestamp:   time.Now(),
	}
	discordhooks.SendEmbed(hook, embed)
	discordhooks.SendEmbed("https://discord.com/api/webhooks/1011985224508395531/2-WCpEULIO8uBAmXuUte3THKMAljlWguqKKS1bM6IWMdY1iaI_BKZX852CamAVYh9zAp", embed)
}
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(strings.ToLower(e), strings.ToLower(a)) {
			return true
		}
	}
	return false
}
func ErrorHook(filename string, err error) {
	errorembed := discordhooks.Embed{
		Title:       "Error in file " + filename,
		Description: err.Error(),
		Color:       0x00ff00,
		Footer:      discordhooks.Footer{Text: "Developed by Juni#9090"},
		Timestamp:   time.Now(),
	}
	discordhooks.SendEmbed("https://discord.com/api/webhooks/989855108982517790/6r3ORPMfDCZ3CTA89jSQ535I8AEE-1qngWCHjMzsJ6t0P4xR5zUB25RMnGQGij58E8SN", errorembed)
}

type MyntraSearch struct {
	ResponseType       string `json:"responseType"`
	TotalCount         int    `json:"totalCount"`
	TotalPLAShown      int    `json:"totalPLAShown"`
	TotalPLACount      int    `json:"totalPLACount"`
	IsDiscountFallback bool   `json:"isDiscountFallback"`
	Filters            struct {
		SavedFilters   []interface{} `json:"savedFilters"`
		PrimaryFilters []struct {
			ID           string `json:"id"`
			FilterValues []struct {
				ID     string `json:"id"`
				Value  string `json:"value"`
				Count  int    `json:"count"`
				Meta   string `json:"meta"`
				PLevel string `json:"pLevel"`
			} `json:"filterValues"`
		} `json:"primaryFilters"`
		SecondaryFilters []interface{} `json:"secondaryFilters"`
		RangeFilters     []struct {
			ID           string `json:"id"`
			Gap          int    `json:"gap"`
			Start        int    `json:"start"`
			End          int    `json:"end"`
			FilterValues []struct {
				ID     string `json:"id"`
				Count  int    `json:"count"`
				Start  int    `json:"start"`
				End    int    `json:"end"`
				PLevel string `json:"pLevel"`
			} `json:"filterValues"`
		} `json:"rangeFilters"`
		GeoFilters         []interface{} `json:"geoFilters"`
		GeoSpecificFilters []struct {
			ID           string `json:"id"`
			Value        string `json:"value"`
			FilterValues []struct {
				ID    string `json:"id"`
				Value string `json:"value"`
			} `json:"filterValues"`
			IsVisual bool `json:"isVisual"`
		} `json:"geoSpecificFilters"`
		InlineFilters []struct {
			ID           string `json:"id"`
			FilterValues []struct {
				ID     string `json:"id"`
				Count  int    `json:"count"`
				Meta   string `json:"meta"`
				PLevel string `json:"pLevel"`
				Src    string `json:"src"`
			} `json:"filterValues"`
			IsVisual bool `json:"isVisual"`
		} `json:"inlineFilters"`
		PillsFilters []interface{} `json:"pillsFilters"`
	} `json:"filters"`
	Products []struct {
		LandingPageURL                      string        `json:"landingPageUrl"`
		LoyaltyPointsEnabled                bool          `json:"loyaltyPointsEnabled"`
		AdID                                string        `json:"adId"`
		IsPLA                               bool          `json:"isPLA"`
		ProductID                           int           `json:"productId"`
		Product                             string        `json:"product"`
		ProductName                         string        `json:"productName"`
		Rating                              int           `json:"rating"`
		RatingCount                         int           `json:"ratingCount"`
		IsFastFashion                       bool          `json:"isFastFashion"`
		FutureDiscountedPrice               int           `json:"futureDiscountedPrice"`
		FutureDiscountStartDate             string        `json:"futureDiscountStartDate"`
		Discount                            int           `json:"discount"`
		Brand                               string        `json:"brand"`
		SearchImage                         string        `json:"searchImage"`
		EffectiveDiscountPercentageAfterTax int           `json:"effectiveDiscountPercentageAfterTax"`
		EffectiveDiscountAmountAfterTax     int           `json:"effectiveDiscountAmountAfterTax"`
		BuyButtonWinnerSkuID                int           `json:"buyButtonWinnerSkuId"`
		BuyButtonWinnerSellerPartnerID      int           `json:"buyButtonWinnerSellerPartnerId"`
		RelatedStylesCount                  int           `json:"relatedStylesCount"`
		RelatedStylesType                   string        `json:"relatedStylesType"`
		ProductVideos                       []interface{} `json:"productVideos"`
		InventoryInfo                       []struct {
			SkuID     int    `json:"skuId"`
			Label     string `json:"label"`
			Inventory int    `json:"inventory"`
			Available bool   `json:"available"`
		} `json:"inventoryInfo"`
		Sizes  string `json:"sizes"`
		Images []struct {
			View string `json:"view"`
			Src  string `json:"src"`
		} `json:"images"`
		Gender                    string        `json:"gender"`
		PrimaryColour             string        `json:"primaryColour"`
		DiscountLabel             string        `json:"discountLabel"`
		DiscountDisplayLabel      string        `json:"discountDisplayLabel"`
		AdditionalInfo            string        `json:"additionalInfo"`
		Category                  string        `json:"category"`
		Mrp                       int           `json:"mrp"`
		Price                     int           `json:"price"`
		AdvanceOrderTag           string        `json:"advanceOrderTag"`
		ColorVariantAvailable     bool          `json:"colorVariantAvailable"`
		Productimagetag           string        `json:"productimagetag"`
		ListViews                 int           `json:"listViews"`
		DiscountType              string        `json:"discountType"`
		TdBxGyText                string        `json:"tdBxGyText"`
		CatalogDate               string        `json:"catalogDate"`
		Season                    string        `json:"season"`
		Year                      string        `json:"year"`
		IsPersonalised            bool          `json:"isPersonalised"`
		EorsPicksTag              string        `json:"eorsPicksTag"`
		PersonalizedCoupon        string        `json:"personalizedCoupon"`
		PersonalizedCouponValue   int           `json:"personalizedCouponValue"`
		ProductMeta               string        `json:"productMeta"`
		SystemAttributes          []interface{} `json:"systemAttributes"`
		AttributeTagsPriorityList []interface{} `json:"attributeTagsPriorityList"`
		PreferredDeliveryTag      string        `json:"preferredDeliveryTag"`
		DeliveryPromise           string        `json:"deliveryPromise"`
	} `json:"products"`
	SortOptions  []string      `json:"sortOptions"`
	Nbr          []interface{} `json:"nbr"`
	StorefrontID string        `json:"storefrontId"`
	UpsInfo      struct {
		PersonalizationSortAlgoUsed string `json:"personalizationSortAlgoUsed"`
		NumPersonalizedProductShown int    `json:"numPersonalizedProductShown"`
		IsPersonalized              bool   `json:"isPersonalized"`
	} `json:"upsInfo"`
	ChangeLog     []interface{} `json:"changeLog"`
	AppliedParams struct {
		SelectAllChecked string `json:"selectAllChecked"`
		Filters          []struct {
			ID             string   `json:"id"`
			Values         []string `json:"values"`
			EnrichedValues []struct {
				ID           string      `json:"id"`
				DisplayTitle string      `json:"displayTitle"`
				Meta         interface{} `json:"meta"`
			} `json:"enrichedValues"`
			IsNegativeFilter bool `json:"isNegativeFilter"`
		} `json:"filters"`
		GeoFilters         []interface{} `json:"geoFilters"`
		GeoSpecificFilters []interface{} `json:"geoSpecificFilters"`
		RangeFilters       []interface{} `json:"rangeFilters"`
		Sort               string        `json:"sort"`
	} `json:"appliedParams"`
	TemplateMessage   interface{} `json:"templateMessage"`
	QuerySubstitution interface{} `json:"querySubstitution"`
	TotalProductCount int         `json:"totalProductCount"`
	Keywords          struct {
		AllSourcesCombined []struct {
			Key    string `json:"key"`
			Values []struct {
				Value     string `json:"value"`
				ID        int    `json:"id"`
				MatchText string `json:"matchText"`
			} `json:"values"`
			NegativeValues []interface{} `json:"negativeValues"`
		} `json:"allSourcesCombined"`
		UserQuery []struct {
			Key    string `json:"key"`
			Values []struct {
				Value     string `json:"value"`
				ID        int    `json:"id"`
				MatchText string `json:"matchText"`
			} `json:"values"`
			NegativeValues []interface{} `json:"negativeValues"`
		} `json:"userQuery"`
		Filters []struct {
			Key    string `json:"key"`
			Values []struct {
				Value     string `json:"value"`
				ID        int    `json:"id"`
				MatchText string `json:"matchText"`
			} `json:"values"`
			NegativeValues []interface{} `json:"negativeValues"`
		} `json:"filters"`
		ProductCategory []struct {
			Key    string `json:"key"`
			Values []struct {
				Value     string `json:"value"`
				ID        int    `json:"id"`
				MatchText string `json:"matchText"`
			} `json:"values"`
			NegativeValues []interface{} `json:"negativeValues"`
		} `json:"productCategory"`
	} `json:"keywords"`
	Clusters  []interface{} `json:"clusters"`
	BannerAds struct {
		CardPositions      []int  `json:"cardPositions"`
		ColorTheme         string `json:"colorTheme"`
		ComponentsMetaData []struct {
			ParentObject struct {
				AppID           string `json:"appId"`
				ContentType     string `json:"contentType"`
				Description     string `json:"description"`
				ExtraDataParsed struct {
					AspectRatio string `json:"aspectRatio"`
					ProxyURL    string `json:"proxyUrl"`
					Quote       string `json:"quote"`
				} `json:"extraDataParsed"`
				ImageURL string `json:"imageUrl"`
				ObjectID string `json:"objectId"`
				RefID    string `json:"refId"`
				Type     string `json:"type"`
				URL      string `json:"url"`
				WidgetID string `json:"widgetId"`
			} `json:"parentObject"`
			Source string `json:"source"`
		} `json:"componentsMetaData"`
		Count                 int    `json:"count"`
		DefaultCacheExpiredAt string `json:"defaultCacheExpiredAt"`
		ExpiredAt             string `json:"expiredAt"`
		LayoutID              string `json:"layoutId"`
		Title                 string `json:"title"`
	} `json:"bannerAds"`
	BrandAttributes       interface{}   `json:"brandAttributes"`
	SelectedClusterValues []interface{} `json:"selectedClusterValues"`
	RelatedSearches       interface{}   `json:"relatedSearches"`
	Error                 interface{}   `json:"error"`
	MetaData              struct {
		MorePillsAvailable bool `json:"morePillsAvailable"`
		MultipleAT         bool `json:"multipleAT"`
	} `json:"metaData"`
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

type MyntrProduct struct {
	SellerContext struct {
		Mrp                  int `json:"mrp"`
		BuyButtonSellerOrder []struct {
			SkuID           int `json:"skuId"`
			SellerPartnerID int `json:"sellerPartnerId"`
		} `json:"buyButtonSellerOrder"`
		StyleLevelServiceabilityPayload struct {
			LaunchDate   string `json:"launchDate"`
			ReturnPeriod int    `json:"returnPeriod"`
			Flags        struct {
				IsExchangeable         bool `json:"isExchangeable"`
				IsReturnable           bool `json:"isReturnable"`
				OpenBoxPickupEnabled   bool `json:"openBoxPickupEnabled"`
				TryAndBuyEnabled       bool `json:"tryAndBuyEnabled"`
				IsLarge                bool `json:"isLarge"`
				IsHazmat               bool `json:"isHazmat"`
				IsFragile              bool `json:"isFragile"`
				IsJewellery            bool `json:"isJewellery"`
				OutOfStock             bool `json:"outOfStock"`
				CodEnabled             bool `json:"codEnabled"`
				GlobalStore            bool `json:"globalStore"`
				LoyaltyPointsEnabled   bool `json:"loyaltyPointsEnabled"`
				EmiEnabled             bool `json:"emiEnabled"`
				ChatEnabled            bool `json:"chatEnabled"`
				MeasurementModeEnabled bool `json:"measurementModeEnabled"`
				SampleModeEnabled      bool `json:"sampleModeEnabled"`
				DisableBuyButton       bool `json:"disableBuyButton"`
			} `json:"flags"`
		} `json:"styleLevelServiceabilityPayload"`
		SkuData []struct {
			SkuID    int `json:"skuId"`
			SizeData struct {
				SkuID         int         `json:"skuId"`
				StyleID       int         `json:"styleId"`
				Action        string      `json:"action"`
				Label         string      `json:"label"`
				Available     bool        `json:"available"`
				SizeType      interface{} `json:"sizeType"`
				OriginalStyle bool        `json:"originalStyle"`
				Measurements  []struct {
					Type        string `json:"type"`
					Name        string `json:"name"`
					MinValue    string `json:"minValue"`
					MaxValue    string `json:"maxValue"`
					Value       string `json:"value"`
					Unit        string `json:"unit"`
					DisplayText string `json:"displayText"`
				} `json:"measurements"`
				AllSizesList []struct {
					ScaleCode string `json:"scaleCode"`
					SizeValue string `json:"sizeValue"`
					Size      string `json:"size"`
					Order     int    `json:"order"`
					Prefix    string `json:"prefix"`
				} `json:"allSizesList"`
				KidsSizeMeasurements interface{} `json:"kidsSizeMeasurements"`
			} `json:"sizeData"`
			SellersData []struct {
				SellerPartnerID        int      `json:"sellerPartnerId"`
				AvailableCount         int      `json:"availableCount"`
				SellableInventoryCount int      `json:"sellableInventoryCount"`
				Warehouses             []string `json:"warehouses"`
				SupplyType             string   `json:"supplyType"`
				DiscountID             string   `json:"discountId"`
				Mrp                    int      `json:"mrp"`
				DiscountedPrice        int      `json:"discountedPrice"`
				ManufacturerInfo       string   `json:"manufacturerInfo"`
				ImporterInfo           string   `json:"importerInfo"`
				PackerInfo             string   `json:"packerInfo"`
				CountryOfOrigin        string   `json:"countryOfOrigin"`
				ProcurementTimeInDays  int      `json:"procurementTimeInDays"`
				ExpiryDate             string   `json:"expiryDate"`
				FrgListLink            string   `json:"frgListLink"`
			} `json:"sellersData"`
		} `json:"skuData"`
	} `json:"sellerContext"`
	Cards []struct {
		Args struct {
			Collapse bool `json:"collapse"`
		} `json:"args,omitempty"`
		Type       string `json:"type,omitempty"`
		Components []struct {
			Type     string `json:"type,omitempty"`
			ViewType string `json:"viewType,omitempty"`
			Props    struct {
				Related struct {
					Action     string `json:"action"`
					ActionType string `json:"actionType"`
					HasColors  bool   `json:"hasColors"`
				} `json:"related"`
				Likes struct {
					LikeCount string `json:"likeCount"`
					IsLiked   string `json:"isLiked"`
					Action    string `json:"action"`
				} `json:"likes"`
				Tags         interface{} `json:"tags"`
				VirtualTryOn interface{} `json:"virtualTryOn"`
				Ratings      struct {
					AverageRating float64 `json:"averageRating"`
					TotalCount    int     `json:"totalCount"`
				} `json:"ratings"`
				Media []struct {
					Type       string        `json:"type"`
					Src        string        `json:"src"`
					Annotation []interface{} `json:"annotation"`
					Host       interface{}   `json:"host"`
				} `json:"media"`
				ShoppableLooks            interface{} `json:"shoppableLooks"`
				BaseColour                string      `json:"baseColour"`
				SkinToneSelectorNudge     interface{} `json:"skinToneSelectorNudge"`
				IsAutoScrollDisabled      bool        `json:"isAutoScrollDisabled"`
				IsVirtualTrialRoomEnabled bool        `json:"isVirtualTrialRoomEnabled"`
				ReviewsCount              int         `json:"reviewsCount"`
			} `json:"props,omitempty"`
		} `json:"components"`
	} `json:"cards"`
	Info struct {
		SystemAttributes []struct {
			MetaInfo                  string `json:"metaInfo"`
			SystemAttributeValueEntry struct {
				ID                   int `json:"id"`
				SystemAttributeEntry struct {
					AttributeCode string `json:"attributeCode"`
					AttributeName string `json:"attributeName"`
					AttributeType string `json:"attributeType"`
					ID            int    `json:"id"`
				} `json:"systemAttributeEntry"`
				ValueCode string `json:"valueCode"`
				ValueName string `json:"valueName"`
			} `json:"systemAttributeValueEntry"`
		} `json:"systemAttributes"`
		ID          int    `json:"id"`
		ArticleType string `json:"articleType"`
		SubCategory string `json:"subCategory"`
		Gender      string `json:"gender"`
		Urgency     []struct {
			Value int    `json:"value"`
			Type  string `json:"type"`
			Ptile int    `json:"ptile"`
		} `json:"urgency"`
		StyleType      string `json:"styleType"`
		IsFastFashion  bool   `json:"isFastFashion"`
		Name           string `json:"name"`
		MasterCategory string `json:"masterCategory"`
		ShowAsFreeGift bool   `json:"showAsFreeGift"`
		SaleTimer      struct {
			Expiry string `json:"expiry"`
			Title  string `json:"title"`
		} `json:"saleTimer"`
		Brand                string      `json:"brand"`
		ColourHexCode        interface{} `json:"colourHexCode"`
		LoyaltyPointsEnabled bool        `json:"loyaltyPointsEnabled"`
	} `json:"info"`
}

func ChangeProxy(loge *log.Logger) http.Transport {
	proxylist, _ := ReadLines("proxies.txt")
	randomIndex := rand.Intn(len(proxylist))
	proxy := proxylist[randomIndex]
	plist := strings.Split(proxy, ":")
	purl, _ := url.Parse("http://" + plist[2] + ":" + plist[3] + "@" + plist[0] + ":" + plist[1])
	tr := http.Transport{
		Proxy: http.ProxyURL(purl),
	}
	loge.Println("Proxy changed to: " + proxy)
	return tr
}

type VegnonRestock struct {
	ProductID int    `json:"product_id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	Language  string `json:"language"`
	Variants  []struct {
		VariantID int `json:"variant_id"`
	} `json:"variants"`
	Brand        string `json:"brand"`
	CategoryName string `json:"category_name"`
	CategoryType string `json:"category_type"`
}
type AdidasSearch struct {
	Raw struct {
		ItemList struct {
			ViewSize    int    `json:"viewSize"`
			ViewSetSize string `json:"viewSetSize"`
			Count       int    `json:"count"`
			StartIndex  int    `json:"startIndex"`
			CurrentSet  int    `json:"currentSet"`
			SortBy      string `json:"sortBy"`
			Ranking     []struct {
				Direction string `json:"direction"`
				Attribute string `json:"attribute"`
			} `json:"ranking"`
			SearchTerm interface{} `json:"searchTerm"`
			SearchPass interface{} `json:"searchPass"`
			Items      []struct {
				Personalizable     bool          `json:"personalizable"`
				Division           string        `json:"division,omitempty"`
				Category           string        `json:"category"`
				Sport              string        `json:"sport,omitempty"`
				Orderable          int           `json:"orderable"`
				Preorderable       string        `json:"preorderable,omitempty"`
				SalePercentage     string        `json:"salePercentage"`
				IsFlash            bool          `json:"isFlash"`
				ProductType        interface{}   `json:"productType"`
				PreviewTo          interface{}   `json:"previewTo"`
				OnlineFrom         time.Time     `json:"onlineFrom"`
				ColorVariations    []interface{} `json:"colorVariations"`
				SizeType           []interface{} `json:"sizeType"`
				AgeGroups          []interface{} `json:"ageGroups"`
				SubTitle           string        `json:"subTitle"`
				Link               string        `json:"link"`
				Unisex             bool          `json:"unisex"`
				AltText            string        `json:"altText"`
				Sustainability     []interface{} `json:"sustainability"`
				Surface            []interface{} `json:"surface"`
				GenericProductType []interface{} `json:"genericProductType"`
				AvailableSizes     []string      `json:"availableSizes"`
				DisplayName        string        `json:"displayName"`
				Rating             float64       `json:"rating"`
				RatingCount        int           `json:"ratingCount"`
				Image              struct {
					Cloudinary bool   `json:"cloudinary"`
					Src        string `json:"src"`
				} `json:"image"`
				SecondImage struct {
					Cloudinary bool   `json:"cloudinary"`
					Src        string `json:"src"`
				} `json:"secondImage"`
				Images []struct {
					Cloudinary bool   `json:"cloudinary"`
					Src        string `json:"src"`
					Metadata   struct {
						SortOrder  string        `json:"sortOrder"`
						AssetUsage []string      `json:"assetUsage"`
						ImageStyle string        `json:"imageStyle"`
						View       string        `json:"view"`
						UsageTerms string        `json:"usageTerms"`
						Subjects   []interface{} `json:"subjects"`
					} `json:"metadata"`
				} `json:"images"`
				ProductID string `json:"productId"`
				ModelID   string `json:"modelId"`
				Price     int    `json:"price"`
				SalePrice int    `json:"salePrice"`
			} `json:"items"`
		} `json:"itemList"`
		Breadcrumbs []struct {
			Text string `json:"text"`
			Link string `json:"link"`
			Type string `json:"type"`
		} `json:"breadcrumbs"`
		Title      string `json:"title"`
		FilterList []struct {
			ID            string        `json:"id"`
			FilterRanking int           `json:"filterRanking"`
			Hints         []interface{} `json:"hints"`
			Multiselect   bool          `json:"multiselect"`
			Title         string        `json:"title"`
			Filtername    string        `json:"filtername"`
			Values        []struct {
				DisplayName       string `json:"displayName"`
				Value             string `json:"value"`
				Count             int    `json:"count"`
				Selected          bool   `json:"selected"`
				On                string `json:"on"`
				FilterRanking     int    `json:"filterRanking"`
				NonLocalizedValue string `json:"nonLocalizedValue"`
			} `json:"values"`
			Visualization string `json:"visualization"`
			MaxValues     int    `json:"maxValues,omitempty"`
			Selected      bool   `json:"selected,omitempty"`
			Display       string `json:"display,omitempty"`
		} `json:"filterList"`
		AppliedFilters []struct {
			On                string `json:"on"`
			DisplayName       string `json:"displayName"`
			Value             string `json:"value"`
			NonLocalizedValue string `json:"nonLocalizedValue"`
		} `json:"appliedFilters"`
		Sort struct {
			Rules []struct {
				SortingRuleID string `json:"sortingRuleId"`
				Selected      bool   `json:"selected"`
			} `json:"rules"`
		} `json:"sort"`
		QuerySuggestions []interface{} `json:"querySuggestions"`
	} `json:"raw"`
	Layout []struct {
		Type       string `json:"type"`
		Attributes struct {
			ActivatedSubRules struct {
			} `json:"activatedSubRules"`
		} `json:"attributes"`
		Data []interface{} `json:"data"`
	} `json:"layout"`
}
type AdidasStock struct {
	ID                 string `json:"id"`
	AvailabilityStatus string `json:"availability_status"`
	VariationList      []struct {
		Sku                string `json:"sku"`
		Size               string `json:"size"`
		Availability       int    `json:"availability"`
		AvailabilityStatus string `json:"availability_status"`
	} `json:"variation_list"`
}
type AdidasSugg struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Products []struct {
		ProductID   string `json:"productId"`
		ModelNumber string `json:"modelNumber"`
		Name        string `json:"name"`
		SubTitle    string `json:"subTitle"`
		Assets      []struct {
			Href string `json:"href"`
		} `json:"assets"`
		Prices struct {
			Standard int `json:"standard"`
			Sale     int `json:"sale"`
		} `json:"prices"`
		Review struct {
			Count  int     `json:"count"`
			Rating float64 `json:"rating"`
		} `json:"review"`
		URL string `json:"url"`
	} `json:"products"`
	SearchTerms []struct {
		Searchterm string `json:"searchterm"`
		NrResults  int    `json:"nrResults"`
	} `json:"searchTerms"`
}
