
require 'sinatra'
require 'net/http'
require 'json'

get '/' do
  # 명언을 가져올 외부 API URL
  url = URI("https://type.fit/api/quotes")

  # API에 GET 요청 보내기
  response = Net::HTTP.get(url)

  # 응답을 JSON으로 파싱
  quotes = JSON.parse(response)

  # 명언 중에서 랜덤으로 선택
  random_quote = quotes.sample

  # 선택된 명언 및 작가 반환
  @quote = random_quote['text']
  @author = random_quote['author']

  erb :index
end
