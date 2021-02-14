package utils

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"math"
	"net/http"
	"page-ss/src/config"
	"page-ss/src/service/logger"
	"strings"
)


func GetFullScreenImageBytes(url string,quality int64) (error, []byte) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(url, quality, &buf)); err != nil {
		log.Fatal(err)
		return err,nil
	}
	return nil,buf
}


func GetFullScreenImageBytesWithHeader(url string,quality int64,headers map[string] interface{} ,cookies []*http.Cookie) (error, []byte) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshotWithHeader(url, quality,headers,cookies, &buf)); err != nil {
		log.Fatal(err)
		return err,nil
	}
	return nil,buf
}



//func GetFullScreenImage(url string,quality int64,outPath string) {
//	// create context
//	ctx, cancel := chromedp.NewContext(context.Background())
//	defer cancel()
//
//	// capture screenshot of an element
//	var buf []byte
//	//if err := chromedp.Run(ctx, elementScreenshot(`https://www.baidu.com/`, `#main`, &buf)); err != nil {
//	//	log.Fatal(err)
//	//}
//	//if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
//	//	log.Fatal(err)
//	//}
//	// capture entire browser viewport, returning png with quality=90
//	if err := chromedp.Run(ctx, fullScreenshot(url, quality, &buf)); err != nil {
//		log.Fatal(err)
//	}
//	if err := ioutil.WriteFile(outPath, buf, 0o644); err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
//}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		//chromedp.WaitVisible(sel, chromedp.BySearch(chromedp.Selector{})),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}





// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshotWithHeader(urlstr string, quality int64,headers map[string]interface{}, cookies []*http.Cookie,res *[]byte) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// add cookies to chrome
		    domain := strings.Replace(config.Global.ProxyUrl, "http://", "", 1 )
			domain = strings.Replace(config.Global.ProxyUrl, "https://", "", 1 )
			for i := 0; i < len(cookies); i += 1 {
				//logger.Log.Info("cookies:",cookies[i])
				//logger.Log.Info("Domain:",cookies[i].Domain)
				//logger.Log.Info("Path:",cookies[i].Path)
				expr := cdp.TimeSinceEpoch(cookies[i].Expires)
				err := network.SetCookie(cookies[i].Name, cookies[i].Value).
					WithExpires(&expr).
					WithDomain(domain).
					WithPath("/").
					WithHTTPOnly(cookies[i].HttpOnly).
					Do(ctx)
				if err != nil {
					logger.Log.Error("Cookie error ",cookies[i],err)
					return err
				}
			}
			return nil
		}),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

