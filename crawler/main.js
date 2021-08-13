const puppeteer = require('puppeteer')
const fetch = require("node-fetch")

const url = "http://127.0.0.1:8080/api/v1/update/shibor"

puppeteer.launch({ headless: true }).then(async browser => {
  const page = await browser.newPage()
  let origin = {}

  try {
    page.on('response', async response => {
      origin = await response.json()
    })
  } catch (err) {
    console.log(err.message)
  }

  await page.goto('http://www.chinamoney.com.cn/r/cms/www/chinamoney/data/shibor/shibor.json?t=1626763368372', { waitUntil: 'load' })

  await browser.close()

  let data = { date: new Date(origin.data.showDateCN).toISOString()}
  origin.records.map(item => {
    data[item.termCode] = +item.shibor
  })

  post(data)
})

async function post(data) {
  try {
    const response = await fetch(url, {
      headers: {
        'Content-Type': 'application/json;charset=utf-8',
      },
      body: JSON.stringify(data),
      method: 'POST',
    })
    console.log(response.statusText)
  } catch (err) {
    console.log(err.message)
    return
  }
}