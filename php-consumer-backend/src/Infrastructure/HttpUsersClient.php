<?php

declare(strict_types=1);

namespace App\Infrastructure;

use GuzzleHttp\Client;
use GuzzleHttp\Psr7\Uri;
use App\Core\UsersClientInterface;
use App\Core\UserDetails;
use function json_decode;

final readonly class HttpUsersClient implements UsersClientInterface
{
    private Client $httpClient;
    public function __construct(
        private string $baseUri,
    ) {
        $this->httpClient = new Client();
    }

    public function getUserDetails(int $id): UserDetails
    {
        $response = $this->httpClient->request('GET', new Uri("{$this->baseUri}/users/{$id}"));
        $body = (string) $response->getBody();
        $data = json_decode($body, true, 512, JSON_THROW_ON_ERROR);

        return new UserDetails($data['id'], $data['email']);
    }
}
