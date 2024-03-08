require 'sinatra'
require 'net/http'
require 'json'

get '/' do
  # Go 서버의 엔드포인트로 HTTP GET 요청을 보냄
  uri = URI.parse('http://localhost:8080/go-endpoint')
  response = Net::HTTP.get_response(uri)
  
  # 응답을 처리
  if response.code == '200'
    @message = JSON.parse(response.body)['message']
  else
    @message = "요청에 실패했습니다."
  end
  
  erb :index
end
