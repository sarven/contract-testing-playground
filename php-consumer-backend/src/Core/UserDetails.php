<?php

declare(strict_types=1);

namespace App\Core;

final readonly class UserDetails
{
    public function __construct(
        public int $id,
        public string $email,
    ) {
    }
}
