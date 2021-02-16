package utils

import (
	"context"
	"errors"
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
	"time"
)


func GetFullScreenImageBytesWithHeader(url string, quality int64, headers map[string]interface{}, cookies []*http.Cookie,delay int64) (error, []byte) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshotWithHeader(url, quality, headers, cookies,delay, &buf)); err != nil {
		log.Fatal(err)
		return err, nil
	}
	return nil, buf
}

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
func fullScreenshotWithHeader(urlstr string, quality int64, headers map[string]interface{}, cookies []*http.Cookie,delay int64, res *[]byte) chromedp.Tasks {

	tasks := chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// add cookies to chrome
			domain := strings.Replace(config.Global.ProxyUrl, "http://", "", 1)
			domain = strings.Replace(config.Global.ProxyUrl, "https://", "", 1)
			for i := 0; i < len(cookies); i += 1 {
				expr := cdp.TimeSinceEpoch(cookies[i].Expires)
				err := network.SetCookie(cookies[i].Name, cookies[i].Value).
					WithExpires(&expr).
					WithDomain(domain).
					WithPath("/").
					WithHTTPOnly(cookies[i].HttpOnly).
					Do(ctx)
				if err != nil {
					logger.Log.Error("Cookie error ", cookies[i], err)
					return err
				}
			}
			return nil
		}),
	}
	headerTasks, err := setHeaders(headers)
	if err == nil {
		tasks = append(tasks, headerTasks)
	}
	tasks = append(tasks, chromedp.Tasks{
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

			//logger.Log.Debug("delay start:",delay)
			time.Sleep(time.Duration(delay)*time.Second)
			//logger.Log.Debug("delay end:",delay)

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
	})
	return tasks
}

func setHeaders(headers map[string]interface{}) (chromedp.Tasks, error) {
	if headers != nil {
		return chromedp.Tasks{
			network.SetExtraHTTPHeaders(network.Headers(headers)),
		}, nil
	} else {
		return nil, errors.New("no header inside")
	}
}
